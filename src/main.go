package main

import (
	"fmt"
	"SendMail/Models"
    m "SendMail/Mailer"
)

func main() {

    // read file CSV/JSON , implement Factory

    // fill data to template string , implement Factory

    // send mail to SMTP/REST API, implement Factory

    // build Dockerfile

    // write README.md
    
	// templateFile, _ := f.GetFileReader(f.EMAIL_TEMPLATE)
    templateReader := Models.EmailTemplateReader{}
	emailTemplate, _ := templateReader.Read("./SampleData/email_template.json")

	fmt.Println("---emailTemplate json:", emailTemplate)

    //read customer file
    customerReader := Models.CustomerReader{}
    customers, _ := customerReader.Read("./SampleData/customers.csv")
	fmt.Println("---customers:", customers)

    // send email
    smtpMailer, _ := m.GetMailer(m.ConfigMailer{
        Type: "SMTP",
    })
    smtpMailer.Send()

}
/*

email_template.json
{
    from: "the marketing<mar@example.com>",
    subject: "a new product...",
    mineType: "text/plain",
    body: "Hi {{TITLE}} {{FIRST_NAME}}",

}

customers.csv
TITLE,FIRST_NAME,LAST_NAME,EMAIL
Mr,John,Smith,john@example.com
Mr,John,Smith,john@example.com

application can be run as:

python send_email.py /path/to/email_template.json /path/to/customers.csv /path/to/output_emails/ /path/to/errors.csv

Expected result
- The application must be tested with unit tests using any unit testing frameworks
- The application should be designed so that it can be updated in the future to send
emails via SMTP or REST API without changing the current code base (can add
new code)
- The provided solution should include a README file which describes how to build
and run the application
- If possible, package the application in a Docker image so that it can run without
any dependencies


*/