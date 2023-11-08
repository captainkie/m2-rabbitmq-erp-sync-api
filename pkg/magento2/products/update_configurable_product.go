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
	"strings"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
	"github.com/captainkie/websync-api/types/request"
)

func UpdateConfigurableProduct(product request.AddUpdateProductRequest, sku, tokens, locale string, visibility int) (string, int, error) {
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

	escapedSKU := url.QueryEscape(sku)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/" + strings.ToLower(locale) + "/V1/products/" + escapedSKU
	jsonStr := prepareConfigurableData(product, sku, cleanedToken, locale, visibility)

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
