package Models

import (
	// "strings"
	"testing"
)

func TestEmailTemplateReader(t *testing.T)  {
	path := "./email_template.json"

	templateReader := &EmailTemplateReader{}

	_, err := templateReader.Read(path)

	if err != nil {
		t.Fatal("EmailTemplateReader must be returned")
	}

}