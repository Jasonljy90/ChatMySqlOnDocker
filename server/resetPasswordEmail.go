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

	//Working but showing raw link, no Reset Password button
	plainTextContent := "You have received this email because a password reset request for Chat App account was received. The reset link will only be valid for 30mins. Click on the button to reset your password: \r\n" + link
	message := mail.NewSingleEmailPlainText(from, subject, to, plainTextContent)

	// plainTextContent not showing in email
	//plainTextContent := "You have received this email because a password reset request for Chat App account was received. The reset link will only be valid for 30mins. Click on the button to reset your password: \r\n"
	//htmlContent := "<a href= " + link + ">Reset Password</a>"
	//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Works but not nice
	//htmlContent := "You have received this email because a password reset request for Chat App account was received. The reset link will only be valid for 30mins. Click on the button to reset your password: \r\n" + "<a href= " + link + ">Reset Password</a>"
	//message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

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
