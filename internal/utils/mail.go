package utils

import (
	"fmt"
	"net/smtp"

	"github.com/CabIsMe/tttn-wine-be/internal"
)

func MailSender(toMail string) {

	// Receiver email address.
	to := []string{
		toMail,
	}

	// smtp server configuration.

	// Message.
	message := []byte("This is a test email message.")
	config := internal.MailEnvs
	// Authentication.
	auth := smtp.PlainAuth("", config.MaiFrom, config.MailPassword, config.SMTPHost)

	// Sending email.
	err := smtp.SendMail(config.SMTPHost+":"+config.SMTPPort, auth, config.MaiFrom, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
