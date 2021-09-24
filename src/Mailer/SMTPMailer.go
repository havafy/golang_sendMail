package Mailer

import (
	"fmt"
)
type SMTPMailer struct {

}

func (m SMTPMailer) Send() {
	fmt.Println("SMTPMailer Send()" )
}