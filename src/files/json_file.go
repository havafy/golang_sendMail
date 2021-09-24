package files

import (
	// "errors"
	"fmt"
)

type JSONReader struct {
	aaa string
}
func (f JSONReader) Read(path string) string{
	f.aaa ="2222"
	fmt.Println("File JSON: " + path)
	return path
}
