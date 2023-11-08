package magentoServiceProduct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func CreateSimpleProduct(product request.AddUpdateProductRequest, tokens string, visibility int) (string, int, error) {
	// get token
	var cleanedToken string
	if tokens == "" {
		token, err := magentoServiceAuth.GetAdminToken()
		if err != nil {
			return "ERROR, Can't Connect to Magento Store API", 500, err
		}

		cleanedToken = helpers.ReplaceAllQuot(token)
	} else {
		cleanedToken = helpers.ReplaceAllQuot(tokens)
	}

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products"
	jsonStr := prepareSimpleData(product, cleanedToken, "EN", visibility, "")

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "ERROR, Can't not connect to m2 store service", 500, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cleanedToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service", 400, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service", 400, err
	}

	if resp.StatusCode != 200 {
		var result map[string]interface{}
		json.Unmarshal([]byte(responseBody), &result)
		msg := fmt.Sprintf("%s", result["message"])
		errSimple := errors.New(msg)

		return "nil", resp.StatusCode, errSimple
	}

	// update thai store
	update, status, errUpdate := UpdateSimpleProduct(product, cleanedToken, "TH", visibility, "not-update-stock")
	if errUpdate != nil {
		fmt.Println("errUpdate: ", errUpdate)
		fmt.Println("update: ", status, update)
	}

	return string(responseBody), resp.StatusCode, nil
}

func prepareSimpleData(product request.AddUpdateProductRequest, token, locale string, visibility int, stock string) []byte {
	// prepare product attribute
	productAttr := prepareSimpleAttribute(product, token, locale)
	// prepare main data
	var status string
	if product.PIC_FILE == "" {
		status = "2"
	} else {
		status = "1"
	}

	var qty string
	if stock == "" || stock == "not-update-stock" {
		qty = product.STOCK_QTY
	} else {
		qty = stock
	}

	var stockItem string
	if stock != "not-update-stock" {
		stockItem = `,
        "stock_item": {
        "qty": ` + qty + `,
        "is_in_stock": true,
        "manage_stock": true, 
        "use_config_manage_stock": true,
        "min_qty": 0,
        "use_config_min_qty": true,
        "min_sale_qty": 1,
        "use_config_min_sale_qty": 1,
        "max_sale_qty": 10000,
        "use_config_max_sale_qty": true,
        "is_qty_decimal": false,
        "backorders": 1,
        "use_config_backorders": true,
        "notify_stock_qty": 1,
        "use_config_notify_stock_qty": true
      },`
	}

	jsonStr := []byte(`{"product": {
    "sku": "` + product.PROD_CODE + `",
    "name": "` + helpers.ReplaceAllSpecialChar(product.PDNAME_EN) + `",
    "attribute_set_id": ` + strconv.Itoa(helpers.AttributeSetID) + `,
    "price": ` + product.UNIT_PRICE + `,
    "status": ` + status + `,
    "visibility": ` + strconv.Itoa(visibility) + `,
    "type_id": "simple",
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
      ]` + stockItem + `
    },
    "custom_attributes": ` + helpers.MappingCustomAttr(product, productAttr, locale, token) + `
  },"saveOptions": true}`)

	return jsonStr
}

func prepareSimpleAttribute(product request.AddUpdateProductRequest, token, locale string) string {
	var attributesEN = map[string]string{
		"product_group":  product.PDGRP_DESC,
		"product_typeof": product.PDTYPE_DESC,
		"product_design": product.PDSIZE_TH,
		"color":          product.COLOR1,
		"size":           product.PDSIZE_DESC,
	}

	var attributesTH = map[string]string{
		"product_group":  product.PDGRP_TH,
		"product_typeof": product.PDTYPE_TH,
		"product_design": product.PDDSGN_TH,
		"color":          product.PDCOLOR_TH,
		"size":           product.PDSIZE_DESC,
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

	slugify := slug.Make(product.PDNAME_EN + "-" + product.PROD_CODE)

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
      },`
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
      },`
	}

	return productAttr
}
