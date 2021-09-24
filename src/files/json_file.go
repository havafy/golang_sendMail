package files

import (
	// "errors"
	"fmt"
	"encoding/json"
)
type TemplateItems struct {
	items []TemplateItem `json:"items"`
}
 
type TemplateItem struct {
	From string `json: "form"`
	Subject string `json: "subject"`
	MineType string `json: "mineType"`
	Body string `json: "body"`
}

type JSONReader struct {
}
func (f JSONReader) Read(path string) string{
	content := fileLoader(path)
	data := TemplateItems{}
 
	_ = json.Unmarshal([]byte(content), &data)
 
	for i := 0; i < len(data.TemplateItems); i++ {
		fmt.Println("From: ", data.TemplateItem[i].From)
		fmt.Println("Subject: ", data.TemplateItem[i].Subject)
	}
	return content
}
