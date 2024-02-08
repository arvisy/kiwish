package helpers

import (
	"fmt"
	"ms-user/model"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func sendMail(email, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "kiwish@kiwish.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(
		os.Getenv("SMTP_SERVER"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendRegisterInfo(email string, User model.User) error {
	subject := "Kiwish user registration"
	content := fmt.Sprintf(`<h1>User successfully registered:</h1>
				<h3>User Info</h3>
				<p>Name: %s</p>
				<p>Email: %s</p>`,
		User.Name,
		User.Email)

	err := sendMail(email, subject, content)
	if err != nil {
		return err
	}
	return nil
}
