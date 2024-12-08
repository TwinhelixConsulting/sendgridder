package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = os.Getenv("DOMAIN") // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = os.Getenv("API_KEY")

func main() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	api_region := os.Getenv("API_REGION")
	if api_region == "" {
		log.Fatal("API_REGION is not set, please set it to either 'us' or 'eu'")
	}

	if api_region == "us" {
		mg.SetAPIBase(mailgun.APIBaseUS)
	} else if api_region == "eu" {
		mg.SetAPIBase(mailgun.APIBaseEU)
	}
	//When you have an EU-domain, you must specify the endpoint:
	// mg.SetAPIBase(mailgun.APIBaseEU)
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	sender := os.Getenv("SENDER")   // e.g.
	subject := os.Getenv("SUBJECT") // e.g. "Hello"
	// read body from file at /tmp/body.txt
	body, err := os.ReadFile("/tmp/body.txt")
	if err != nil {
		log.Fatal(err)
	}

	// body := "Hello from Mailgun Go!"
	recipient := os.Getenv("RECIPIENT")

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, string(body), recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
