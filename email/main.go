package main

import (
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

func main() {
	abc := gomail.NewMessage()

	email := os.Getenv("smtp_email")
	secret := os.Getenv("smtp_secret")

	abc.SetHeader("From", email)
	abc.SetHeader("To", email)
	abc.SetHeader("Subject", "Enter message subject here")
	abc.SetBody("text/plain", "Enter body here")

	a := gomail.NewDialer("smtp.fastmail.com", 587, email, secret)

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
