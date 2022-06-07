package main

import "fmt"

func main() {

	// Penggunakan map
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	delete(chicken, "januari")

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
