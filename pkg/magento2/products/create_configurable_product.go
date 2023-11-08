package magentoServiceProduct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAttribute "github.com/captainkie/websync-api/pkg/magento2/attributes"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
	"github.com/captainkie/websync-api/types/payload"
	"github.com/captainkie/websync-api/types/request"
	"github.com/gosimple/slug"
)

func CreateConfigurableProduct(product request.AddUpdateProductRequest, tokens string) (string, string, int, error) {
	// get token
	var cleanedToken string
	var createMapping string
	if tokens == "" {
		token, err := magentoServiceAuth.GetAdminToken()
		if err != nil {
			return "ERROR, Can't Connect to Magento Store API", "", 500, err
		}

		cleanedToken = helpers.ReplaceAllQuot(token)
	} else {
		cleanedToken = helpers.ReplaceAllQuot(tokens)
	}

	// pretty print product
	// productByte, err := json.Marshal(product)
	// if err != nil {
	// 	return "ERROR, Can't Convert Product to Byte", err
	// }
	// helpers.PrintPrettyJson(productByte)

	// Substring SKU 8 Digit To master configurable product
	// sku := product.PROD_CODE[0:8]
	sku := helpers.PadString(product.PROD_CODE, 8, '#')
	sku = strings.ReplaceAll(sku, "#", "")

	// Step 1. Create simple products
	visibility := 1
	createSimple, statusCode, errCreateSimple := CreateSimpleProduct(product, cleanedToken, visibility)
	if errCreateSimple != nil {
		return "ERROR, Can't not create simple product", "", statusCode, errCreateSimple
	}

	// Step 2. Create master configurable product
	createConfigurable, statusCode, errFindConfig := GetProductBySKU(cleanedToken, sku)
	if errFindConfig != nil {
		// Not Found Master, Then create simple first
		fmt.Printf("errFindConfig SKU : %s => CODE=%d, MSG=%s", sku, statusCode, errFindConfig)

		if statusCode == 404 {
			// create master configurable product
			visibility := 4
			serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products"
			jsonStr := prepareConfigurableData(product, sku, cleanedToken, "EN", visibility)

			req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonStr))
			if err != nil {
				return "ERROR, Can't not connect to m2 store service", "", 500, err
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+cleanedToken)
			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				return "ERROR, Can't not read response body from m2 store service", "", 400, err
			}

			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "ERROR, Can't not read response body from m2 store service", "", 400, err
			}

			if resp.StatusCode != 200 {
				var result map[string]interface{}
				json.Unmarshal([]byte(responseBody), &result)
				msg := fmt.Sprintf("%s", result["message"])
				errConfig := errors.New(msg)

				return "nil", "", resp.StatusCode, errConfig
			}

			// update thai store
			update, status, errUpdate := UpdateConfigurableProduct(product, sku, cleanedToken, "TH", visibility)
			if errUpdate != nil {
				fmt.Println("errUpdate: ", errUpdate)
				fmt.Println("update: ", status, update)
			}

			// set this simple product is master configurable product
			createMapping = product.PROD_CODE

			createConfigurable = string(responseBody)
		}
	}

	// Step 3. Assign simple products to master configurable product
	// Define configurable product options color, size
	defineConfigurableAttribute(sku, "93", "Color", "0", cleanedToken)
	defineConfigurableAttribute(sku, "144", "Size", "1", cleanedToken)

	// Link the simple products to the configurable product
	var simplePayload payload.SimpleProductPayload
	errPayload := json.Unmarshal([]byte(createSimple), &simplePayload)
	if errPayload != nil {
		return "ERROR, Can't not get simple product sku", "", 400, errPayload
	}

	linkConfig, err := linkSimpleToConfigurable(sku, simplePayload.Sku, cleanedToken)
	if err != nil {
		fmt.Println("err link: ", err)
		fmt.Println("linkConfig: ", linkConfig)
	}

	return string(createConfigurable), createMapping, 200, nil
}

func prepareConfigurableData(product request.AddUpdateProductRequest, sku, token, locale string, visibility int) []byte {
	// prepare product attribute
	productAttr := prepareConfigurableAttribute(product, sku, token, locale)
	// prepare main data
	var status string
	if product.PIC_FILE == "" {
		status = "2"
	} else {
		status = "1"
	}

	jsonStr := []byte(`{"product": {
    "sku": "` + sku + `",
    "name": "` + helpers.ReplaceAllSpecialChar(product.PDNAME_EN) + `",
    "attribute_set_id": ` + strconv.Itoa(helpers.AttributeSetID) + `,
    "price": 0,
    "status": ` + status + `,
    "visibility": ` + strconv.Itoa(visibility) + `,
    "type_id": "configurable",
    "weight": ` + strconv.Itoa(helpers.DefaultWeight) + `,
    "extension_attributes": {
      "website_ids": [
        ` + strconv.Itoa(helpers.WebsiteID) + `
      ],
      "category_links": [
        {
          "position": 0,
          "category_id": "` + helpers.MappingCategory(product.PDBRND_CODE, product.PDGRP_CODE) + `"
        },
        {
          "position": 1,
          "category_id": "` + helpers.MappingFindInShopCategory() + `"
        },
        {
          "position": 2,
          "category_id": "` + helpers.MappingSubCategory(product.PDTYPE_CODE) + `"
        }
      ],
      "stock_item": {
        "is_in_stock": true,
        "manage_stock": true,
        "use_config_manage_stock": true
      }
    },
    "custom_attributes": ` + helpers.MappingConfigurableCustomAttr(product, productAttr, locale, token) + `
  },"saveOptions": true}`)

	return jsonStr
}

func prepareConfigurableAttribute(product request.AddUpdateProductRequest, sku, token, locale string) string {
	var attributesEN = map[string]string{
		"product_group":  product.PDGRP_DESC,
		"product_typeof": product.PDTYPE_DESC,
		"product_design": product.PDSIZE_TH,
	}

	var attributesTH = map[string]string{
		"product_group":  product.PDGRP_TH,
		"product_typeof": product.PDTYPE_TH,
		"product_design": product.PDDSGN_TH,
	}

	// loop attributes
	var productAttr string = ``
	for key, val := range attributesEN {
		// get attribute options id
		getOption, statusCode, err := magentoServiceAttribute.GetAttributeOptionByCode(token, "all", key)
		if err != nil {
			fmt.Println("err: ", err)
			fmt.Println("statusCode: ", statusCode)
		}

		var productAttributeOptionPayload payload.ProductAttributeOptionPayload
		err = json.Unmarshal([]byte(getOption), &productAttributeOptionPayload)
		if err != nil {
			fmt.Println("err: ", err)
		}

		// loop to get attribute option id
		var productGroupID string
		for _, v := range productAttributeOptionPayload {
			if strings.TrimSpace(strings.ToLower(v.Label)) == strings.TrimSpace(strings.ToLower(val)) {
				productGroupID = v.Value
			}
		}

		if productGroupID == "" && val != "" {
			// add new option
			productGroupID, err = magentoServiceAttribute.CreateAttributeOption(token, key, attributesEN[key], attributesTH[key])
			if err != nil {
				fmt.Println("err: ", err)
			}

			productGroupID = helpers.ReplaceAllQuot(productGroupID)
		}

		if productGroupID != "" {
			productAttr += `
      {
        "attribute_code": "` + key + `",
        "value": "` + productGroupID + `"
      },`
		}
	}

	slugify := slug.Make(product.PDNAME_EN + "-" + sku)

	productAttr += `
      {
        "attribute_code": "product_model",
        "value": "` + product.PDMODEL_DESC + `"
      },
      {
        "attribute_code": "product_weight",
        "value": "` + product.WEIGHT + `"
      },
      {
        "attribute_code": "country_of_manufacture",
        "value": "TH"
      },
      {
        "attribute_code": "url_key",
        "value": "` + slugify + `"
      },`

	if locale == "EN" {
		productAttr += `
      {
        "attribute_code": "product_material",
        "value": "` + product.MATERIAL1 + `"
      },
      {
        "attribute_code": "product_dimension",
        "value": "` + helpers.ReplaceAllSpecialChar(product.DIMENSION_DESC_EN) + `"
      },
      {
        "attribute_code": "product_uom",
        "value": "` + product.UOM_EN + `"
      }
    `
	} else {
		productAttr += `
      {
        "attribute_code": "product_material",
        "value": "` + product.MATERIAL1_TH + `"
      },
      {
        "attribute_code": "product_dimension",
        "value": "` + helpers.ReplaceAllSpecialChar(product.DIMENSION_DESC_TH) + `"
      },
      {
        "attribute_code": "product_uom",
        "value": "` + product.UOM_TH + `"
      }
    `
	}

	return productAttr
}

func defineConfigurableAttribute(sku, id, label, index, token string) {
	escapedSKU := url.QueryEscape(sku)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/configurable-products/" + escapedSKU + "/options"

	jsonStr := []byte(`{"option": {
    "attribute_id": ` + id + `,
    "label": "` + label + `",
    "position": 0,
    "is_use_default": 0,
    "values": [
      {
        "value_index": ` + index + `
      }
    ]
  }}`)

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("err: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err: ", err)
	}

	defer resp.Body.Close()
}

func linkSimpleToConfigurable(sku, childSku, token string) (string, error) {
	escapedSKU := url.QueryEscape(sku)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/configurable-products/" + escapedSKU + "/child"

	jsonStr := []byte(`{"childSku": "` + childSku + `"}`)

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "ERROR, Can't not connect to m2 store service", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service", err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service", err
	}

	return string(responseBody), nil
}
