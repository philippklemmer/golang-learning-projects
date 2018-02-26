/*
	--- Udemy Course ---
	This is "Hands On 2" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/001_prereq
*/
package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}
func (p Person) pSpeak() {
	fmt.Println(p.name, `says, "Wassup"!`)
}

type Agent struct {
	Person
	secret bool
}
func (a Agent) saSpeak() {
	fmt.Println(a.name, `says, "Good Morning!"`)
}


func main () {
	fmt.Println("Hello to Preperation 2 of Learn Webdevelopment with Golang")

	luca := Person{
		name: "Luca Doinet",
		age: 22,
	}

	mrBond := Agent {
		Person{
			name: "James Bond",
			age: 36,
		},
		true,
	}

	//format
	fmt.Printf("%s is %d years old\n", luca.name, luca.age)

	//polymorphism
	fmt.Printf("The Agent %q, which is %d years old. \nAgentStatus: %t \n", 
		mrBond.Person.name, 
		mrBond.Person.age, 
		mrBond.secret,
	)

	// using object methods 
	luca.pSpeak()
	mrBond.saSpeak()
}