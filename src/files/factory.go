package files

import (
	"errors"
	"fmt"
)

const (
	CSV = 1
	JSON = 2
)
type FileReader interface {
	Read(path string) string
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
