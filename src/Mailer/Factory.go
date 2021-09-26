package Mailer

import (
	"errors"
	"fmt"
	"SendMail/Models"
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
type Mailer interface{
	Send(customers []Models.Customer, template Models.EmailTemplateItem) error
}

// Implement Factory Pattern
func GetMailer(c ConfigMailer) (Mailer, error) {
	switch c.Type {
		case "SMTP":
			return &SMTPMailer{config: c}, nil
		case "API":
			return &APIMailer{config: c}, nil
		default:
		return nil, errors.New(fmt.Sprintf("Mailer type %d not recognized.", c.Type))
	}
}
