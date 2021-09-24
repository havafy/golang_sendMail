package Mailer

import (
	"errors"
	"fmt"
)

type ConfigMailer struct {
	Type string
	ApiKey string
	Host string
	Port string
	User string
	Password string
}
type Mailer interface{
	Send()
}

// Factory Pattern
func GetMailer(c ConfigMailer) (Mailer, error) {
	switch c.Type {
		case "SMTP":
			return new(SMTPMailer), nil
		case "API":
			return new(APIMailer), nil
		default:
		return nil, errors.New(fmt.Sprintf("Mailer type %d not recognized.", c.Type))
	}
}
