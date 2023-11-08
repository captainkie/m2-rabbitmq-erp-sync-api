package magentoServiceAttribute

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
)

func GetAttribute(tokens string) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/attributes?searchCriteria"

	req, err := http.NewRequest("GET", serviceURL, nil)
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
		errAuth := errors.New(msg)

		return "nil", resp.StatusCode, errAuth
	}

	return string(responseBody), resp.StatusCode, nil
}

func GetAttributeByCode(tokens, code string) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/attributes/" + code

	req, err := http.NewRequest("GET", serviceURL, nil)
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
		errAuth := errors.New(msg)

		return "nil", resp.StatusCode, errAuth
	}

	return string(responseBody), resp.StatusCode, nil
}

func GetAttributeOptionByCode(tokens, locale, code string) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/" + locale + "/V1/products/attributes/" + code + "/options"

	req, err := http.NewRequest("GET", serviceURL, nil)
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
		errAuth := errors.New(msg)

		return "nil", resp.StatusCode, errAuth
	}

	return string(responseBody), resp.StatusCode, nil
}
