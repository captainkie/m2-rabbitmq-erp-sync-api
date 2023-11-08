package magentoServiceMedia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
)

type GetMediaQuery struct {
	EditMode    bool `json:"editMode"`
	ForceReload bool `json:"forceReload"`
}

func createMagentoRequestWithAuth(method, url, token string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func getMediaDataFromMagento(tokens, serviceURL string) (string, int, error) {
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

	req, err := createMagentoRequestWithAuth("GET", serviceURL, cleanedToken)
	if err != nil {
		return "ERROR, Can't not connect to M2 store service", 500, err
	}

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
		errAuth := errors.New(msg)
		return "nil", resp.StatusCode, errAuth
	}

	return string(responseBody), resp.StatusCode, nil
}

func GetMediaBySKU(tokens, sku string) (string, int, error) {
	escapedSKU := url.QueryEscape(sku)
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/" + escapedSKU + "/media"

	return getMediaDataFromMagento(tokens, serviceURL)
}
