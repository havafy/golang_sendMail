package Mailer

import (
	"SendMail/Models"
	"errors"
	"net/smtp"
)

func newSMTPMailer(c ConfigMailer) IMailer {
	return &SMTPMailer{
		config: c,
	}
}

type SMTPMailer struct {
	customers []Models.Customer
	template  Models.EmailTemplate
	config    ConfigMailer
}

func (m *SMTPMailer) SetCustomers(customers []Models.Customer) {
	m.customers = customers
}

func (m *SMTPMailer) SetTemplate(template Models.EmailTemplate) {
	m.template = template
}
func (m *SMTPMailer) SendAll() error {
	if m.config.Host == "" {
		return errors.New("Missing SMTP config.")
	}
	for _, customer := range m.customers {

		// fill customer info to email template
		contentBody := fillCustomerToTemplate(m.template.Body, customer)

		// send email by SMTP protocol
		m.Send(customer, m.template.Subject, contentBody)

	}
	return nil
}

func (m *SMTPMailer) Send(customer Models.Customer, subject string, contentBody string) error {
	config := m.config
	// toList is list of email address that email is to be sent.
	toList := []string{customer.Email}
	// This is the message to send in the mail
	msg := []byte("From: " + config.From + "\r\n" +
		"To: " + customer.Email + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		contentBody + "\r\n")
	body := []byte(msg)
	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

	err := smtp.SendMail(config.Host+":"+config.Port, auth, config.From, toList, body)

	// handling the errors
	if err != nil {
		return err
	}
	return nil
}
