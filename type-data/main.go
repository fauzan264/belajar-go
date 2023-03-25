package main

import (
	"fmt"
	"reflect"
)

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Println(reflect.TypeOf(negativeNumber))

	// decimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // menghasilkan tipe data float dengan 3 digit dibelakang koma

	// tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// tipe data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var newMessage = `Nama saya "Fauzan".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(newMessage)

	/*
		nil bukan merupakan tipe data, melainkan sebuah nilai. Variable yang isi nilainya nil berarti memiliki nilai kosong.
		Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal,
		tetap akan ada nilai default-nya.
		- Zero value dari string adalah "" (string kosong)
		- Zero value dari bool adalah false
		- Zero value dari tipe numerik non-desimal adalah 0
		- Zero value dari tipe numerik desimal adalah 0.0
	*/

	/*
		Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. nil tidak bisa digunakan
		pada tipe data yang sudah dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya
		dengan nil, di antaranya:
		- pointer
		- tipe data fungsi
		- slice
		- map
		- channel
		- interface kosong atau any (yang merupakan alias dari interface{} )
	*/
}