package main

import "fmt"

func main() {

	// Penerapan 1
	var namaPertama [4]string
	namaPertama[0] = "trafalgar"
	namaPertama[1] = "d"
	namaPertama[2] = "water"
	namaPertama[3] = "law"

	fmt.Println(namaPertama[0], namaPertama[1], namaPertama[2], namaPertama[3])

	// Pengisian element array yang melebihi alokasi awal
	var namaKedua [4]string
	namaKedua[0] = "trafalgar"
	namaKedua[1] = "d"
	namaKedua[2] = "water"
	namaKedua[3] = "law"
	// baris kode dibawah ini menghasilkan error
	namaKedua[4] = "ez"

	// Inisialisasi nilai awal array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruits = [4]string{"apple", "grepe", "banana", "melon"}

	// Perulangan elemen array menggunakan for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// Perulangan element array menggunakan for - range
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// Perulangan element array menggunakan for - range dengan variable _(underscore)
	for _, fruit := range fruits {
		fmt.Printf("nama buah %s\n", fruit)
	}
}
