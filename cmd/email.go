package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	smtpHost := "smtp.rambler.ru"
	smtpPort := "587"
	smtpUser := "ariocp@rambler.ru"
	smtpPass := "ImIs7QsZDI"

	to := "ariocp@yandex.ru" // адрес получателя
	subject := "Subject: Test Email\n"
	body := "This is a test email."
	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{to}, msg)
	if err != nil {
		log.Fatalf("smtp error: %s", err)
	}

	fmt.Println("Email sent successfully")
}
