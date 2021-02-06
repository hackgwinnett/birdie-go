package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
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

func send() {

	var t string
	fmt.Println("Is this message for sponsors or members? ")
	fmt.Scanln(&t)
	if t == "sponsors" {

		var to []string
		var msg []byte

		file, err := os.Open("email.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		b, err := ioutil.ReadAll(file)
		fmt.Print(b)
		msg = b

		file2, err := os.Open("temp.txt") // read sponsors email list
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()

		scanner := bufio.NewScanner(file2)
		for scanner.Scan() {

			from := "hackgwinnett@gmail.com"
			password := "hackgwinn@489$"
			to = []string{scanner.Text()}

			smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

			// Authentication.
			auth := smtp.PlainAuth("", from, password, smtpServer.host)

			// Sending email.
			err := smtp.SendMail(smtpServer.Address(), auth, from, to, msg)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if t == "members" {
		var to []string
		var msg []byte

		file, err := os.Open("email.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		b, err := ioutil.ReadAll(file)
		fmt.Print(b)
		msg = b

		file2, err := os.Open("temp.txt") // read sponsors email list
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()

		scanner := bufio.NewScanner(file2)
		for scanner.Scan() {

			from := "XXXX"
			password := "XXXX"
			to = []string{scanner.Text()}

			smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

			// Authentication.
			auth := smtp.PlainAuth("", from, password, smtpServer.host)

			// Sending email.
			err := smtp.SendMail(smtpServer.Address(), auth, from, to, msg)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Emails sent successfully")
	}

}
