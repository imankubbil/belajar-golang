package main

import "fmt"

func main() {

	var fruits = []string{"apple", "grape", "banana", "melon"}

	// Inisialisasi slice
	// fmt.Println(fruits[0]) // "apple"

	/* var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Printf("fruitsA : %s, fruitsB : %s, fruitsC : %c", fruitsA, fruitsB, fruitsC) */

	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits) // ["apple", "grape"]

	// Fungsi len
	// fmt.Println(len(fruits)) // 4

	// Fungsi cap
	// fmt.Println(len(fruits)) // len: 4
	// fmt.Println(cap(fruits)) // cap: 4

	// var aFruits = fruits[0:3]
	// fmt.Println(len(aFruits)) // len: 3
	// fmt.Println(cap(aFruits)) // cap: 4

	// var bFruits = fruits[1:4]
	// fmt.Println(len(bFruits)) // len: 3
	// fmt.Println(cap(bFruits)) // cap: 3

	// Fungsi append
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "banana", "melon"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "melon", "papaya"]

	// Fungsi copy
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2
}
