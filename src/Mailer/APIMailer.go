package Mailer

import (
	"fmt"
	"SendMail/Models"
)
type APIMailer struct {
	config ConfigMailer
}

func (m APIMailer) Send(customers []Models.Customer, template Models.EmailTemplateItem) error{
	// config := m.config
	for _, customer := range customers {
		// ... todo something with send by REST API
		toList := []string{customer.Email}
		fmt.Println("APIMailer sent: ", toList)
	  }

	  return nil
  }