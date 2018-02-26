/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.3" of the learning
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
	// slice - composite literal - directly passing data to it
	sages := []string{
		"Gandhi",
		"MLK",
		"Buddha",
		"Jesus",
		"Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}