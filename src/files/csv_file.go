package files

import (
//	"errors"
	"fmt"
)

type CSVReader struct {}
func (f *CSVReader) Read(path string) string{
	fmt.Println("File CSV: " + path)
	return path
}
