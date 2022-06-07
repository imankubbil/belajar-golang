package main

import "fmt"

func main() {

	/*
		Tipe data numerik non decimal
	*/
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	/*
		Tipe data numerik decimal
	*/

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	/*
		Tipe data boolean
	*/

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	/*
		Tipe data string
	*/

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
}
