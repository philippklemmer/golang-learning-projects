/*
Going trough the Udemy Course of 'Webdevelopment with Go.
Doing the first exercises in this file.
This only used for improving Golang Language purposes.
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
	sq := Square{20,20}
	rt := Reactangle{30,20} 
	
	info(sq)
	info(rt)
}






