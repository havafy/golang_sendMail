package Mailer

import (
  "fmt"
	"SendMail/Models"
	"net/smtp"
)
type SMTPMailer struct {
	config ConfigMailer
}

func (m SMTPMailer) Send(customers []Models.Customer, template Models.EmailTemplateItem) (bool, error){
	config := m.config
  for _, customer := range customers {
      // toList is list of email address that email is to be sent.
      toList := []string{customer.Email}
  
      // This is the message to send in the mail
      msg := []byte("From: "+ config.From +"\r\n" +
          "To: " + customer.Email + "\r\n" +
          "Subject: " + template.Subject + "\r\n\r\n" +
          template.Body + "\r\n")
      body := []byte(msg)
      auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

      err := smtp.SendMail(config.Host+":"+config.Port, auth, config.From, toList, body)

      // handling the errors
      if err != nil {
        return false, err
      }
      fmt.Println("sent: ", customer.Email)
    }
	return true, nil
}