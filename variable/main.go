package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"


	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")


	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstNamee = "ahmad"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastNamee := "fauzan"
	
	fmt.Println(firstNamee, lastNamee)

	// multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first + " ", second + " ", third)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println(seventh + " ", eight + " ", ninth)

	// berbeda-beda tipe data
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one)
	fmt.Println(isFriday)
	fmt.Println(twoPointTwo)
	fmt.Println(say)

	// variable underscore
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "Juang", "Fauzan"
	fmt.Println(name)

	// deklarasi variable menggunakan keyword new
	// new digunakan untuk membuat variabel pointer dengan tipe data tertentu
	name1 := new(string)
	fmt.Println(name1) // menampung data bertipe pointer string
	fmt.Println(*name1)

	// keyword make
	// keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:
	// - channel
	// - map
	// - slice
}