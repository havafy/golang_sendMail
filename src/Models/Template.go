package Models

import (
	// "errors"
	"os"
	"encoding/json"
	"io/ioutil"
)

type EmailTemplateItem struct {
	From string `json: "from"`
	Subject string `json: "subject"`
	MineType string `json: "mineType"`
	Body string `json: "body"`
}

type EmailTemplateReader struct {
}
func (f EmailTemplateReader) Read(path string) (EmailTemplateItem, error){
	var results EmailTemplateItem

    file, err := os.Open(path)
	if err != nil {
		return results, err
	}
	defer file.Close()
    byteValue, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(byteValue, &results)

	return results, nil
}
