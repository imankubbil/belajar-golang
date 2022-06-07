package main

import "fmt"

func main() {
	var firstName string = "Nur"

	var lastName string
	lastName = "iman"

	/*
		manifest typing / bisa disebut juga dengan binding
	*/
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	var personTwo string = "john"
	afterPersonTwo := "wick"

	fmt.Printf("halo %s %s!\n", personTwo, afterPersonTwo)
}
