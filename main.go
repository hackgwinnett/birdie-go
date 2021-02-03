package main

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server

type smtpServer struct {
	host string
	port string
}

// Address URL to smtp server

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	// Sender data.
	from := "XXXX"
	password := "XXXX"

	// Receiever email addresses
	// Change to parsing from a .txt file in the future
	to := []string{
		"XXXX",
		"XXXX",
	}

	// smtp server configuration
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	// Message.
	message := []byte("XXXX")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
