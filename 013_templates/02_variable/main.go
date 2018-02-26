/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.1" of the learning
	webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}