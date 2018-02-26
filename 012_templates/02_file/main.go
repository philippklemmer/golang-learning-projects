/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 2.01" of the learning webdevelopment with GoLang
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

func main() {
	fmt.Println("Welcome to learning templating. #part2.1")
	
	// parse files
	tpl, err := template.ParseFiles("../tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//create new file and close it if func main ends
	nf, err := os.Create("./index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// using equal sign here because err is already defined
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}