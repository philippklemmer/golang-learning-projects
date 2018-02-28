/*
--- Udemy Course ---
	This is "Learn about Templates in Go Part 4.3" of the learning
	webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/
*/
package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"tl":      strings.ToLower,
	"fdouble": double,
	"fsquare": square,
	"fsqrt":   sqRoot,
}

func init() {
	tpl = template.Must(template.New("data.gohtml").Funcs(fm).ParseFiles("data.gohtml"))
}

func double(x float64) float64 {
	return float64(x) + float64(x)
}

func square(x float64) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.Execute(os.Stdout, 3.0)
	if err != nil {
		log.Fatalln(err)
	}
}
