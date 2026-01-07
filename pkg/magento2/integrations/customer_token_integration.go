package magentoServiceAuth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
)

func GetCustomerToken() (string, error) {
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/V1/integration/customer/token"

	params := map[string]string{
		"username": os.Getenv("MAGE_SERVICE_USER"),
		"password": os.Getenv("MAGE_SERVICE_PASS"),
	}

	jsonPayload, err := json.Marshal(params)
	helpers.ErrorPanic(err)

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "ERROR, Can't not connect to m2 store service authenticator ", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service authenticator", err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR, Can't not read response body from m2 store service authenticator", err
	}

	if resp.StatusCode != 200 {
		var result map[string]interface{}
		json.Unmarshal([]byte(responseBody), &result)
		msg := fmt.Sprintf("%s", result["message"])
		errAuth := errors.New(msg)

		return "", errAuth
	}

	return string(responseBody), nil
}
