package Mailer

import (
	"fmt"
)
type APIMailer struct {

}

func (m APIMailer) Send() {
	fmt.Println("APIMailer Send()")
}