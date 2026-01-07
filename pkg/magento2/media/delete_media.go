package magentoServiceMedia

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
)

type DeleteDirectoryRequest struct {
	Directory string `json:"directory"`
}
type DeleteMediaRequest struct {
	Path string `json:"path"`
}

func DeleteDirectory(tokens, path string) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/V1/image/delete-directory"

	requestData := DeleteDirectoryRequest{
		Directory: path,
	}

	fmt.Println("DeleteDirectory requestData", requestData)

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("DeleteDirectory jsonData", string(jsonData))

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "ERROR, Can't not connect to m2 store service", 500, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cleanedToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR, Can't not read response body from M2 store service", 400, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR, Can't not read response body from M2 store service", 400, err
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

func DeleteMedia(tokens, path string) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/V1/image/delete-image"

	requestData := DeleteMediaRequest{
		Path: path,
	}

	fmt.Println("DeleteMedia requestData", requestData)

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("DeleteMedia jsonData", string(jsonData))

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "ERROR, Can't not connect to m2 store service", 500, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cleanedToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR, Can't not read response body from M2 store service", 400, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR, Can't not read response body from M2 store service", 400, err
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
