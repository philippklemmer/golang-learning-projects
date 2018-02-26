/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 2.0" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main () {
	fmt.Println("Welcome to learning templating. #part2.0")	

	tpl, err := template.ParseFiles("../tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}