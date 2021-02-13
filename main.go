package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

type smtpServer struct {
	host string
	port string
}

func main() {
	var count int32
	fmt.Println("-send: Send your emails")
	fmt.Println("-store: Store your sponsors/members emails")
	fmt.Println("-terminate: end the program")
	fmt.Println("-help: know the commands")
	fmt.Println("-inst: add your credentials to the database")


	for {

		 var choice string
		 if count >= 1 {
		 	fmt.Println("\n")
		 	fmt.Println("-")
		 	fmt.Scanln(&choice)


		 } else {
			 fmt.Println("Choice? ")
			 fmt.Scanln(&choice)
		 }




		if choice == "send" {
			send()
		}
		if choice == "terminate" {
			fmt.Println("!")
			os.Exit(3)
		}
		if choice == "store" {
			store()
		}

		if choice == "help" {
			help()
		}

		if choice == "inst" {
			inst()
		}

		count += 1


	}




}

func help() {
	fmt.Println("-send: Send your emails")
	fmt.Println("-store: Store your sponsors/members emails")
	fmt.Println("-terminate: end the program")
}



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
		var user string
		var pwd string

		file, err := os.Open("email.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		file3, err := os.Open("creds.txt")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scann := bufio.NewScanner(file3)
		scann.Split(bufio.ScanLines)
		var txtlines []string

		for scann.Scan() {
			txtlines = append(txtlines, scann.Text())
		}

		file3.Close()

		user = txtlines[0]
		pwd = txtlines[1]


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

			from := user
			password := pwd
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
		var user string
		var pwd string

		file3, err := os.Open("creds.txt")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scann := bufio.NewScanner(file3)
		scann.Split(bufio.ScanLines)
		var txtlines []string

		for scann.Scan() {
			txtlines = append(txtlines, scann.Text())
		}

		file3.Close()

		user = txtlines[0]
		pwd = txtlines[1]


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
		msg = []byte("To: bill@gates.com\r\n" +
			"Subject: Why are you not using Mailtrap yet?\r\n" +
			"\r\n" +
			"Here’s the space for our great sales pitch\r\n")

		file2, err := os.Open("members.txt") // read sponsors email list
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()

		scanner := bufio.NewScanner(file2)
		for scanner.Scan() {

			from := user
			password := pwd
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

func single() {
	// individual
	var t string
	fmt.Println("Is this message for a company or individual?")
	fmt.Scanln(&t)
	if t == "company" { // individual company

		var to []string
		var request string
		var check bool
		var msg []byte

		fmt.Println("Company email? ")
		fmt.Scanln(&request)

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

		file2, err := os.Open("temp.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file2.Close()
		scanner := bufio.NewScanner(file2)
		for scanner.Scan() {
			if scanner.Text() == request {
				check = true
			} else {
				continue
			}
		}

		if check == true {
			// Sender data.
			from := "XXXX"
			password := "XXXX"
			to = []string{request}

			smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

			// Authentication.
			auth := smtp.PlainAuth("", from, password, smtpServer.host)

			// Sending email.
			err := smtp.SendMail(smtpServer.Address(), auth, from, to, msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Email Sent!")
			// Receiever email addresses
			// Change to parsing from a .txt file in the future
		} else {
			log.Fatal("Email requested not found in database. Please add this requested email to the database.")
		}
	}

	if t == "individual" { // individual person
		// same code but modifications for membership emails
	}
}

func store() {

	var option string

	fmt.Println("Enter option (single, multiple (include directory of text file)): ")
	fmt.Scanln(&option)

	if option == "single" {

		fmt.Println("Enter a sponsor name: ")

		var company string

		fmt.Scanln(&company)

		//Write first line
		err := ioutil.WriteFile("temp.txt", []byte(company+"\n"), 0644)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer file.Close()

		//Print the contents of the file
	}

	if option == "multiple" {
		file2, err := os.Open("temp.txt") // read sponsors email list
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()

		scanner := bufio.NewScanner(file2)
		for scanner.Scan() {
			err := ioutil.WriteFile("temp.txt", []byte(scanner.Text()+"\n"), 0644)
			if err != nil {
				log.Fatal(err)
			}

			file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				file.Close()
				log.Println(err)
			}

			// Print the contents of the file

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Elements written successfully")
	}

}

func app(s string) {
	f, err := os.OpenFile("creds.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(s + "\n"); err != nil {
		log.Println(err)
	}
}

func inst() {
	var username string
	var password string

	fmt.Scanln(&username)
	fmt.Scanln(&password)

	app(username)
	app(password)

}

func direct() {
	myfile, e := os.Create("creds.txt")
	if e != nil {
		log.Fatal(e)
	}
	myfile.Close()


	myfile2, e := os.Create("temp.txt")
	if e != nil {
		log.Fatal(e)
	}
	myfile2.Close()


	myfile3, e := os.Create("email.txt")
	if e != nil {
		log.Fatal(e)
	}
	myfile3.Close()
}