package magentoServiceCatgory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetCategories(token string) (string, int, error) {
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/categories"

	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		return "ERROR, Can't not connect to m2 store service", 500, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
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
