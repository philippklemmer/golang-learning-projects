/*
--- Udemy Course ---
	This is "Learn about Templates in Go Part 4.4" of the learning
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

var tpl *template.Template

type person struct {
	Name  string
	Job   string
	Admin bool
}

func init() {
	tpl = template.Must(template.New("data.gohtml").ParseFiles("data.gohtml"))
}

func main() {

	philipp := person{
		Name:  "Philipp",
		Job:   "Go Lover",
		Admin: true,
	}

	jakub := person{
		Name:  "Jakub",
		Job:   "Managmant Student",
		Admin: false,
	}

	kostek := person{
		Name:  "Konstancja",
		Job:   "Mum",
		Admin: false,
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	users := []person{philipp, jakub, kostek}

	data := []interface{}{
		users,
		numbers,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
