package magentoServiceMedia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAuth "github.com/captainkie/websync-api/pkg/magento2/integrations"
	"github.com/captainkie/websync-api/types/request"
)

func UpdateMedia(tokens, sku string, id int, media request.CreateMediaRequest) (string, int, error) {
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

	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/" + sku + "/media/" + fmt.Sprintf("%d", id)

	fmt.Println(serviceURL)

	jsonData, err := json.Marshal(media)
	if err != nil {
		fmt.Println("Error:", err)
	}

	req, err := http.NewRequest("PUT", serviceURL, bytes.NewBuffer(jsonData))
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

	return string(responseBody), resp.StatusCode, nil
}
