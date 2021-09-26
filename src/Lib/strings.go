package Lib

import (
	"strings"
	"time"
)

func FillVarsToTemplate(content string, vars map[string]string) string{
	for key, item := range vars {
		content = strings.ReplaceAll(content, "{{"+key+"}}", item)
	}
	return content
}

func DateFormat(date time.Time) string{
	return date.Format("02 Jan 2006")
}

