/*
	--- Udemy Course ---
	This is "Hands On 1" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/001_prereq
*/
package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

func info(sh Shape) float64 {
	fmt.Println(sh)
	fmt.Println(sh.area())	
	return  sh.area()
}

type Square struct {
	width float64 
	height float64
}

type Reactangle struct {
	length float64
	height float64
}

func (r Reactangle) area() float64 {
	return  r.length * r.height
}

func (s Square) area() float64 {
	return s.width * s.height
}

func main () {
	fmt.Println("Hello to Preperation 1 of Learn Webdevelopment with Golang")
	
	sq := Square{20,20}
	rt := Reactangle{30,20} 
	
	info(sq)
	info(rt)
}






