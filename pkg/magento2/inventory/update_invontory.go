package magentoServiceInventory

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
	"github.com/captainkie/websync-api/types/request"
)

func UpdateInventory(stock request.UpdateStockRequest, tokens string, qty int, is_in_stock bool) (string, int, error) {
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

	escapedSKU := url.QueryEscape(stock.PROD_CODE)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/" + escapedSKU
	jsonStr := prepareInventoryData(stock, qty, is_in_stock)

	req, err := http.NewRequest("PUT", serviceURL, bytes.NewBuffer(jsonStr))
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

	return string(responseBody), resp.StatusCode, nil
}

func UpdateStockItems(payload request.StockItemsRequest, sku, tokens string) (string, int, error) {
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

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
	}

	escapedSKU := url.QueryEscape(sku)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/" + escapedSKU + "/stockItems/" + fmt.Sprintf("%d", payload.StockItem.ItemId)

	req, err := http.NewRequest("PUT", serviceURL, bytes.NewBuffer(jsonData))
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

	return string(responseBody), resp.StatusCode, nil
}

func prepareInventoryData(stock request.UpdateStockRequest, qty int, is_in_stock bool) []byte {
	jsonStr := []byte(`{"product": {
    "extension_attributes": {
      "stock_item": {
        "qty": ` + fmt.Sprintf("%d", qty) + `,
        "is_in_stock": ` + fmt.Sprintf("%t", is_in_stock) + `
      }
    }
  }}`)

	return jsonStr
}
