package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendEmail(email string, firstName string, lastName, link string) {
	from := mail.NewEmail("Live Chat", os.Getenv("SENDER_EMAIL"))
	subject := "Password Reset Request"
	to := mail.NewEmail(firstName+" "+lastName, email)
	plainTextContent := "You have received this email because a password reset request for Chat App account was received. The reset link will only be valid for 30mins. Click the link to reset your password: \r\n" + link
	//htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	message := mail.NewSingleEmailPlainText(from, subject, to, plainTextContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
