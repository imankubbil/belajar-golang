package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	/*
		Variable Temporary pada if - else
	*/
	var pointSecond = 8840.0

	if percent := pointSecond / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	/*
		Kondisi menggunakan switch - case
	*/

	var pointSwitchCase = 7

	switch pointSwitchCase {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*
		Menggunakan switch - case banyak kondisi
	*/

	var pointMoreCondition = 6

	switch pointMoreCondition {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*
		Kurung kurawal pada keyword case & default
	*/

	var pointKurungKurawal = 3

	switch pointKurungKurawal {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	/*
		Kondisi bersarang pada if else dan switch case
	*/

	var lastPoint = 10

	if lastPoint > 7 {
		switch lastPoint {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if lastPoint == 5 {
			fmt.Println("not bad")
		} else if lastPoint == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if lastPoint == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
