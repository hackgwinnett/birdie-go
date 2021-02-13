package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

const colorRed = "\033[0;31m"
const colorGreen = "\033[0;32m"
const colorBlue = "\033[0;34m"
const colorNone = "\033[0m"

type smtpServer struct {
	host string
	port string
}

func main() {


	var count int32
	help()


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
			appEmail()
		}

		if choice == "help" {
			help()
		}

		if choice == "inst" {
			inst()
		}
		if choice == "init" {
			dir()
		}

		count += 1


	}


}

func col(c string, s string) string {
	return c + s + colorNone
}



func help() {
	fmt.Println("-send: Send your emails")
	fmt.Println("-store: Store your sponsors/members emails")
	fmt.Println("-terminate: end the program")
	fmt.Println("-help: know the commands")
	fmt.Println("-inst: add your credentials to the database")
	fmt.Println("-init: initialize the directory")
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

func appEmail() {

	var choice string
	var e string
	var ult string
	fmt.Println("Add to your members or sponsors? ")
	fmt.Scanln(&choice)

	if choice == "members" {
		fmt.Println("Member to add? ")
		fmt.Scanln(&e)
		ult = "members.txt"
		lr, err := os.OpenFile(ult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer lr.Close()
		if _, err := lr.WriteString(e + "\n"); err != nil {
			log.Println(err)

		}
	}
		if choice == "sponsors" {
			fmt.Println("Sponsor to add? ")
			fmt.Scanln(&e)
			ult = "temp.txt"
			lr, err := os.OpenFile(ult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}

			defer lr.Close()
			if _, err := lr.WriteString(e + "\n"); err != nil {
				log.Println(err)

			} else {
				log.Fatal("Choose a valid option: members or sponsors to append a new email.")
			}

		}
		if choice == "members" {
			for {
				var opt string
				fmt.Println("Add another member? (yes or no) ")
				fmt.Scanln(&opt)
				if opt == "yes"{
					fmt.Println("Member to add? ")
					fmt.Scanln(&e)
					ult = "members.txt"
					lr, err := os.OpenFile(ult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Println(err)
					}

					defer lr.Close()
					if _, err := lr.WriteString(e + "\n"); err != nil {
						log.Println(err)

					}
				} else {
					break
				}
			}
		}
			if choice == "sponsors" {
				for {
					var opt string
					fmt.Println("Add another sponsor? (yes or no) ")
					fmt.Scanln(&opt)
					if opt == "yes"{
						fmt.Println("Sponsor to add? ")
						fmt.Scanln(&e)
						ult = "temp.txt"
						lr, err := os.OpenFile(ult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Println(err)
						}

						defer lr.Close()
						if _, err := lr.WriteString(e + "\n"); err != nil {
							log.Println(err)

						}
					} else {
						break
					}
				}
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

func dir() {
	creds, e := os.Create("creds.txt")
	if e != nil {
		log.Fatal(e)
	}
	creds.Close()


	sponsors, e := os.Create("temp.txt")
	if e != nil {
		log.Fatal(e)
	}
	sponsors.Close()


	email, e := os.Create("email.txt")
	if e != nil {
		log.Fatal(e)
	}
	email.Close()

	members, e :=  os.Create("members.txt")
	if e != nil {
		log.Fatal(e)
	}
	members.Close()
}