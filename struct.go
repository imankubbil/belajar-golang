package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 student
	s1.name = "Nur Iman"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("grade :", s1.grade)

	// Inisialisasi Object Struct
	var s2 = student{}
	s2.name = "Abdulloh"
	s2.age = 20
	s2.grade = 2

	var d3 = student{
		grade: 2,
		person: person{
			name: "Abdulloh",
			age:  21,
		},
	}

	var d1 = student{
		grade: 1,
		person: person{
			name: "Sarif",
			age:  20,
		},
	}

	fmt.Println("student 1 :", s2.name)
	fmt.Println("student 2 :", d3.name)
	fmt.Println("student 3 :", d1.name)

	var s3 = student{
		grade: 1,
		person: person{
			name: "wayne",
			age:  17,
		},
	}

	var s4 = student{
		grade: 1,
		person: person{
			name: "bruce",
			age:  18,
		},
	}

	fmt.Println("student 4 :", s3.name)
	fmt.Println("student 5 :", s4.name)
}
