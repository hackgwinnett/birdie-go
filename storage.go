package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var option string

	fmt.Println("Enter option (single, multiple (ordered line by line)): ")
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
		data, err := ioutil.ReadFile("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
