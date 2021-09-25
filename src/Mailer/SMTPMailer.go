package Mailer

import (
	"SendMail/Models"
	"net/smtp"
)
type SMTPMailer struct {
	config ConfigMailer
}

func (m SMTPMailer) Send(data Models.EmailTemplateItem) (bool, error){
	config := m.config

    // toList is list of email address that email is to be sent.
    toList := []string{"test@gmail.com"}
 
    
    // This is the message to send in the mail
    msg := []byte("From: "+ config.From +"\r\n" +
        "To: test@gmail.com\r\n" +
        "Subject: Test mail\r\n\r\n" +
        "Email body\r\n")
    body := []byte(msg)
    auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

    err := smtp.SendMail(config.Host+":"+config.Port, auth, config.From, toList, body)

    // handling the errors
    if err != nil {
		return false, err
    }
	return true, nil
}