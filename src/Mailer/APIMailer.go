package Mailer

import (
	"fmt"
	"SendMail/Models"
)
type APIMailer struct {
	config ConfigMailer
}

func (m APIMailer) Send(customers []Models.Customer, template Models.EmailTemplateItem) (bool, error){
	// config := m.config
	for _, customer := range customers {
		// toList is list of email address that email is to be sent.
		toList := []string{customer.Email}

		fmt.Println("APIMailer sent: ", toList)
	  }
	  
	  return true, nil
  }