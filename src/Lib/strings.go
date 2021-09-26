package Lib
import (
	"fmt"
	"text/template"
	"bytes"

)

func FillVarToTemplate() string{
	//str := "Hi {{TITLE}} {{FIRST_NAME}"
	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = t.Execute( &tpl, map[string]string{ "Name": "222", "Description": "111"})
	if err != nil {
		panic(err)
	}
	fmt.Println("--->", tpl.String())
	return tpl.String()

}
func funcMap(){

}

