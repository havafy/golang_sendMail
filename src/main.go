package main

import (
	"SendMail/Builder"
	"SendMail/Config"
	"SendMail/Mailer"
	"SendMail/Models"
	"fmt"
)

func main() {

	configVars := Config.EnvLoader()
	template := &Models.EmailTemplate{}
	emailTemplate, _ := template.Get("./SampleData/email_template.json")

	// read customer file
	customer := &Models.Customer{}
	customers, _ := customer.Get("./SampleData/customers.csv")

	// report customer missing email address
	export := &Builder.Export{}
	export.SetCustomers(customers)
	export.SetTemplate(emailTemplate)

	// export JSON file to a folder
	export.ToPath("./SampleData/Export")

	// export invalid customer to CSV file
	export.ErrorToPath("./SampleData/error.csv")

	// set config values to Mailer
	mailer, _ := Mailer.GetMailer(Mailer.ConfigMailer{
		Type:      configVars["EMAIL_TYPE"], // set SMTP or REST API mode on .env file
		ApiKey:    configVars["EMAIL_API_KEY"],
		Host:      configVars["EMAIL_HOST"],
		Port:      configVars["EMAIL_PORT"],
		User:      configVars["EMAIL_AUTH_USER"],
		Password:  configVars["EMAIL_AUTH_PASSWORD"],
		From:      configVars["EMAIL_FROM"],
		From_name: configVars["EMAIL_FROM_NAME"],
	})

	// set data from input files
	mailer.SetCustomers(customers)
	mailer.SetTemplate(emailTemplate)

	// start to send with customers and content
	// mailer.SendAll()

	fmt.Println("Sent completed.")
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
