package magentoServiceAttribute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type StoreLabel struct {
	StoreID int    `json:"store_id"`
	Label   string `json:"label"`
}

type Option struct {
	Label       string       `json:"label"`
	Value       string       `json:"value"`
	StoreLabels []StoreLabel `json:"store_labels"`
}

type AttrOptionRequest struct {
	Option `json:"option"`
}

func CreateAttributeOption(token, code, label1, label2 string) (string, error) {
	serviceURL := os.Getenv("MAGE_HOST") + "/rest/all/V1/products/attributes/" + code + "/options"

	// jsonStr := []byte(`{"option": {
	//   "label": "` + strings.TrimSpace(label1) + `",
	//   "value": "` + strings.TrimSpace(label1) + `",
	//   "store_labels": [
	//     {
	//       "store_id": 1,
	//       "label": "` + strings.TrimSpace(label2) + `"
	//     },
	//     {
	//       "store_id": 2,
	//       "label": "` + strings.TrimSpace(label1) + `"
	//     }
	//   ]
	// }}`)

	thLabel := strings.TrimSpace(label2)
	if thLabel == "" {
		thLabel = strings.TrimSpace(label1)
	}

	requestData := AttrOptionRequest{
		Option: Option{
			Label: strings.TrimSpace(strings.ToLower("_" + label1)),
			Value: strings.TrimSpace(label1),
			StoreLabels: []StoreLabel{
				{
					StoreID: 1,
					Label:   thLabel,
				},
				{
					StoreID: 2,
					Label:   strings.TrimSpace(label1),
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonData))
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
