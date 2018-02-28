/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 3.7" of the learning
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
	Name	string
	Motto	string
}

type pokemon struct {
	Name	string
	Color	string
	Tone	string
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

	pika := pokemon{
		Name: 	"Pikachu",
		Color:	"Yellow",
		Tone:	"Pika Pika",
	}

	pidgeon := pokemon{
		Name:	"Pidgeon",
		Color:	"Grey/Black",
		Tone:	"Piep Piep",
	}

	sages := []sage{buddha, gandhi, mlk}
	pokis := []pokemon{pika, pidgeon}

	data := struct {
		Wisdom	[]sage
		Fluffys []pokemon
	}{
		sages,
		pokis,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}