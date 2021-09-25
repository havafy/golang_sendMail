package Mailer

import (
	"fmt"
	"SendMail/Models"
)
type APIMailer struct {
	config ConfigMailer
}

func (m APIMailer) Send(data Models.EmailTemplateItem) (bool, error){
	fmt.Println("APIMailer Send()" )
	return true, nil
}