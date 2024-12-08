package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	sender := os.Getenv("SENDER")
	subject := os.Getenv("SUBJECT")
	recipient := os.Getenv("RECIPIENT")
	apiKey := os.Getenv("API_KEY")
	hasHtml := os.Getenv("HAS_HTML")
	cc := os.Getenv("CC")

	if apiKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}
	from := mail.NewEmail("", sender)
	to := mail.NewEmail("", recipient)

	body, err := os.ReadFile("/tmp/body.txt")
	if err != nil {
		log.Fatal(err)
	}
	message := mail.NewSingleEmail(from, subject, to, string(body), "")

	if hasHtml == "true" {
		html, err := os.ReadFile("/tmp/body.html")
		if err != nil {
			log.Fatal(err)
		}
		message = mail.NewSingleEmail(from, subject, to, string(body), string(html))
	}
	if cc != "" {
		personalization := mail.NewPersonalization()
		personalization.AddTos(mail.NewEmail("", recipient))
		personalization.AddCCs(mail.NewEmail("", cc))
		message.AddPersonalizations(personalization)
	}
	client := sendgrid.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status Code: %d, Response: %s\n", response.StatusCode, response.Body)
}
