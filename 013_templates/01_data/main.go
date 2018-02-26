/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.0" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main () {
	/*
		You can pass data to tpl.Executes last parameter.
		It's aggregate data that you can pass in, that means
		it can be multiple interleaved data passed into the
		template.
		You can access the data through a full stop (.)
	*/
	err := tpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}