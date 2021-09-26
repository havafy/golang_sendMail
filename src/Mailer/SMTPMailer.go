package Mailer

import (
  "fmt"
	"SendMail/Models"
// 	"net/smtp"
  "SendMail/Lib"
  "time"
)
type SMTPMailer struct {
	config ConfigMailer
}

func (m SMTPMailer) Send(customers []Models.Customer, template Models.EmailTemplateItem) error{
//	config := m.config
  for _, customer := range customers {
      // toList is list of email address that email is to be sent.
     // toList := []string{customer.Email}

      vars := map[string]string{ 
        "TITLE": customer.Title, 
        "FIRST_NAME": customer.First_name, 
        "LAST_NAME": customer.Last_name, 
        "TODAY": Lib.DateFormat(time.Now()),
      }
    
      contentBody := Lib.FillVarsToTemplate(template.Body, vars)
      fmt.Println("--->: ", contentBody, vars)
      // This is the message to send in the mail
      // msg := []byte("From: "+ config.From +"\r\n" +
      //     "To: " + customer.Email + "\r\n" +
      //     "Subject: " + template.Subject + "\r\n\r\n" +
      //     contentBody + "\r\n")
      // body := []byte(msg)
      // auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

      // err := smtp.SendMail(config.Host+":"+config.Port, auth, config.From, toList, body)

      // handling the errors
      // if err != nil {
      //   return err
      // }
      // fmt.Println("sent: ", customer.Email)
    }
	return nil
}