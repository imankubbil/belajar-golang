package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// Immediately - Invoked Function Expression (IIFE)
	var numberSecond = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numberSecond {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(2) // Ciri khas IIFE adalah adanya kurung parameter tepat setelah deklarasi closure berakhir. Jika ada parameter

	fmt.Println("original number :", numberSecond)
	fmt.Println("filtered number :", newNumbers)

	// Closure sebagai nilai kembalian
	var maxClosure = 3
	var numberClosure = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numberClosure, maxClosure)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numberClosure)
	fmt.Printf("find \t: %d\n\n", maxClosure)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
