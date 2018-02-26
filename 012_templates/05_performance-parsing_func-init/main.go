/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 2.5" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"os"
	"log"
	"text/template"
)

// tpl is accessible in this package
var tpl *template.Template

// init will call after reading imports and initializing variables
// we only want to parse templates ones #performance
func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}