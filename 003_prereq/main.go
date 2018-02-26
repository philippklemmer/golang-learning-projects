/*
	--- Udemy Course ---
	This is "Hands On 3" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/001_prereq
*/
package main

import (
	"fmt"
)

type Human interface {
	speak()
}

type Person struct {
	fname string
	lname string
}

func (p Person) speak() {
	fmt.Println(p.fname, p.lname, ` says "Good Morning Bitches"`)
}

type SecretAgent struct {
	Person
	license bool
}

func (sa SecretAgent) speak() {
	fmt.Println(sa.fname, sa.lname,` says, "Good Night Bitches"`)
}

func main () {
	fmt.Println("Hello to Preperation 3 of Learn Webdevelopment with Golang")

	phil := Person{
		fname: "Philipp",
		lname: "Klemmer",
	}

	mrBond := SecretAgent {
		Person{
			fname: "James",
			lname: "Bond",
		},
		true,
	}

	phil.speak()
	mrBond.speak()

}