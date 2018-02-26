/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.6" of the learning
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
	Motto	string
}

func init () {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main () {

	buddha := sage{
		Name:	"Buddha",
		Motto: 	"The belief of no beliefs",
	}

	gandhi := sage{
		Name:	"Gandhi",
		Motto:	"Be the change",
	}

	mlk := sage{
		Name:	"Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	sages := []sage{buddha,gandhi,mlk}


	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}