/*
--- Udemy Course ---
	This is "Learn about Templates in Go Part 5.0" of the learning
	webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) Agedouble() int {
	return p.Age * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	philipp := person{
		Name: "philipp",
		Age:  23,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", philipp)
	if err != nil {
		log.Fatalln(err)
	}
}
