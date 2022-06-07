package main

import "fmt"

func main() {

	// Cara 1
	for i := 0; i < 5; i++ {
		fmt.Println("Angka dari i", i)
	}

	// Cara 2
	var j = 0
	for j < 5 {
		fmt.Println("Angka dari j", j)
		j++
	}

	// Cara 3
	var k = 0
	for {
		fmt.Println("Angka dari k", k)

		k++
		if k == 5 {
			break
		}
	}

	// Cara 4 for - range biasa digunakan untuk me-looping data bertipe array

	// Cara 5
	for l := 1; l <= 10; l++ {
		if l%2 == 1 {
			continue
		}

		if l > 8 {
			break
		}

		fmt.Println("Angka dari l", l)
	}

	// Perulangan bersarang
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}

		fmt.Println()
	}

	// Label dalam perulangan
	outerLoop: // <- ini sebuah label
	for o := 0; o < 5; o++ {
		for p := 0; p < 5; p++ {
			if o == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", o, "][", p, "]", "\n")
		}
	}
}
