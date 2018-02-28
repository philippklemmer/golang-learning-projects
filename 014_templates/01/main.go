/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 4.0" of the learning
	webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type person struct {
	Name string
	Job  string
}

func init() {
	tpl = template.Must(template.New("data.gohtml").Funcs(fm).ParseFiles("data.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	pete := person{
		"Pete McMuffin",
		"Supervisor @Google",
	}
	angelo := person{
		"Angelo Joe Killer",
		"Pro Golfer",
	}
	luca := person{
		"Luca Doinet",
		"BWL STUDENT MUAHAHA",
	}

	persons := []person{pete, angelo, luca}

	err := tpl.Execute(os.Stdout, persons)
	if err != nil {
		log.Fatalln(err)
	}
}
