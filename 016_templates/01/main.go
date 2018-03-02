/*
--- Udemy Course ---
	This is "Learn about Templates in Go Part 6.0" of the learning
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

type course struct {
	Number, Name string
	Units        int
}

func (c course) DashToColon() string {
	return strings.Join(strings.Split(c.Number, "-"), ":")
}

type semester struct {
	Term    string
	Courses []course
}

type season struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	season := season{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", 4},
				course{"CSCI-130", "Introduction to Web Programming with Go", 4},
				course{"CSCI-140", "Mobile Apps Using Go", 4},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", 5},
				course{"CSCI-190", "Advanced Web Programming with Go", 5},
				course{"CSCI-191", "Advanced Mobile Apps With Go", 5},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", season)
	if err != nil {
		log.Fatalln(err)
	}
}
