package files

import (
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	CSV = 1
	JSON = 2
)
type FileReader interface {
	Read(path string) string
}
func fileLoader(path string) string{
	b, err := ioutil.ReadFile(path) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    return string(b) // convert content to a 'string'
}

// Using concept of Factory Pattern
func GetFileReader(fileType int) (FileReader, error) {

	switch fileType {
		case CSV:
			return new(CSVReader), nil
		case JSON:
			return new(JSONReader), nil
		default:
		return nil, errors.New(fmt.Sprintf("File type %d not recognized.", fileType))
	}
}
