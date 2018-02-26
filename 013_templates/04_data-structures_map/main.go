/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.4" of the learning
	webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main () {
	// map - composite literal - using string as a key and also as value
	sages := map[string]string{
		"India": 	"Ganush",
		"USA": 		"Captain America",
		"Germany": 	"NOT Merkel",
		"Uganda": 	"Ufafafe",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}