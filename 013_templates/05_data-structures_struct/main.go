/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.5" of the learning
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

type sage struct {
	Name 	string
	Motto 	string
}

func init () {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main () {

	buddha := sage{
		Name:	"Buddha",
		Motto:	"The belief of not beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}