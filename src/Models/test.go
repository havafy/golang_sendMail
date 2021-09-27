package Models

import (
	// "strings"
	"testing"
)

func TestEmailTemplate(t *testing.T)  {
	path := "./email_template.json"

	template := &EmailTemplate{}

	_, err := template.Get(path)

	if err != nil {
		t.Fatal("template must be returned")
	}

}
