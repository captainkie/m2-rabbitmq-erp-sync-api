package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintPrettyJson(data []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(prettyJSON.String())
}

func IsJSONString(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
