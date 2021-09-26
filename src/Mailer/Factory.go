package Mailer

import (
	"errors"
	"fmt"
	"SendMail/Models"
	"SendMail/Lib"
	"time"
)

type ConfigMailer struct {
	Type string
	ApiKey string
	Host string
	Port string
	User string
	Password string
	From string
	From_name string
}
type IMailer interface{
	SetCustomers(customers []Models.Customer)
	SetTemplate(template Models.EmailTemplateItem)
	SendAll() error
	Send(customer Models.Customer, subject string, contentTemplate string) error
}

func fillCustomerToTemplate(template string, customer Models.Customer) string{
	vars := map[string]string{ 
	  "TITLE": customer.Title, 
	  "FIRST_NAME": customer.First_name, 
	  "LAST_NAME": customer.Last_name, 
	  "TODAY": Lib.DateFormat(time.Now()),
	}
	return Lib.FillVarsToTemplate(template, vars)
  }
// Implement Factory Pattern
func GetMailer(c ConfigMailer) (IMailer, error) {
	switch c.Type {
		case "SMTP":
			return &SMTPMailer{config: c}, nil
		case "API":
			return &APIMailer{config: c}, nil
		default:
		return nil, errors.New(fmt.Sprintf("Mailer type %d not recognized.", c.Type))
	}
}
