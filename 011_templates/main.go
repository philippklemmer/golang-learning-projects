/*
	--- Udemy Course ---
	This is "Learn about Templates in Go Part 1" of the learning webdevelopment with GoLang
	from GoesToEleven
	https://github.com/GoesToEleven/golang-web-dev/tree/master/001_prereq
*/
package main

import (
	"fmt"
	"time"
	"os"
	"io"
	"log"
	"strings"
)

type User struct {
	name string
	timestamp string
}

func createDataSet(name string, timestamp string) string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Welcome ` + name +`</title>
		</head>
		<body>
			<h1>Welcome ` + name + `! You logged in @` + timestamp +`
		</body>
		</html>
	`
}

func writeDataToFile(filename string, data string) {
	nf, err := os.Create(filename)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	//always close file!
	defer nf.Close()

	// copy copies from src to destination until EOF is reached on source 
	// or an error occurs
	if _, err := io.Copy(nf, strings.NewReader(data)); err != nil {
		log.Fatal("error while writing to file", err)
	}
}

func main() {

	// get current timestamp as string and set timezone to berlin
	_ , err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println("Couldn't read out local timezones", err)
		return
	}
	t := time.Now().Format("Jan 2")

	// create user for data
	u := User{
		name: "Philipp",
		timestamp: t,
	}

	fmt.Println("Welcome to learning templating. #part1")

	fmt.Println("Creating dataset")
	data := createDataSet(u.name, u.timestamp)
	fmt.Println(data)

	fmt.Println("Write Data to file")
	writeDataToFile("index.html", data)

	fmt.Println("open new file in browser")
	os.exec.Command("firefox", "index.html").Output()

}