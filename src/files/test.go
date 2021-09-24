package files

import (
	"strings"
	"testing"
)

func TestJSONFileReader(t *testing.T)  {
	path := "./template.json"
	file, err := GetFileReader(JSON)

	if err != nil {
		t.Fatal("File type 'JSON' must be returned")
	}

	resultStr := file.Read(path)
	if !strings.Contains(resultStr, "was sent to number") {
		t.Error("JSON Reader result contect was not correct.")
	}
}
