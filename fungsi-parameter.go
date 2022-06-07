package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Nur", "Ahmad", "Abdulloh"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [Nur Ahmad Abdullah]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [Abdulloh]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [Ahmad]
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
