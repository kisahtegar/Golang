package main

import (
	"fmt"
	"math"
)

func main() {
	// ========================================================================================================
	// [Manifest Typing]
	// Go memiliki aturan dalam hal penulisan variable. tipe data yang digunakan harus dituliskan juga.
	// var <nama-variable> <tipe-data>
	var firstName string = "Kisah"
	var lastName string = ""
	lastName = "Abdi"
	var age int = 20
	var alamat string = "Pasar Kemis"

	fmt.Println("Nama saya ", firstName)
	fmt.Println("Umur saya ", age)
	fmt.Printf("Hallo %s %s usia %d \n", firstName, lastName, 20)

	// ========================================================================================================
	// [Inference Typing]
	// Metode deklarasi variable yang tipe data-nya ditentukan oleh tipe data nilainya, cara
	// kontradiktif jika dibandingkan dengan cara sebelumnya.
	// <nama-variable := <value>
	middleName := "Tegar"
	fmt.Println("Nama tengah saya", middleName)

	// ========================================================================================================
	// Study case: Buat variable nama, umur dan alamat.
	fmt.Printf("Nama saya %s umur saya %d alamat %s \n", firstName, age, alamat)

	// ========================================================================================================
	// [Multiple Variable]
	// Menuliskan variable-varible dengan pembatas tanda koma (,).
	var first, second, third string
	first, second, third = "Satu", "Dua", "Tiga"
	fmt.Printf("Hallo %s %s %s \n", first, second, third)

	// ========================================================================================================
	// [Variable Underscore]
	// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak
	// dipakai. Bisa dibilang variable ini merupakan keranjang sampah.
	organisasi, _ := "HMSI", "UNIPI"
	fmt.Println(organisasi)

	// ========================================================================================================
	// [TIPE DATA]

	// 1) Tipe Data Numerik Non-Desimal
	// var positiveNumber uint8 = 255
	// var negativeNumber = -1243424644

	// 2) Tipe Data Numerik Desimal
	// Ada 2 tipe data numerik desimal pada GO yaitu float32 & float64
	// Float32
	// 	- Reprensentasi bilang real 32-bit
	// 	- Rentang nilai sekitar 10^-38 hingga 10^38

	// Float64
	// 	- Reprensentasi bilang real 64-bit
	// 	- Rentang nilai sekitar 10^-308 hingga 10^308

	var myDesimal float64 = 36.89
	fmt.Printf("Bilangan desimal: %f\n", math.Round(myDesimal))

	// 3) Tipe Data Boolean
	var isNegative bool = false
	fmt.Printf("boolean : %t \n", isNegative)

	// 4) Tipe Data String
	// Selain dengan petik dua (""). String juga dapat ditulis dengan tanda backtip (``).
	var pesan = `Nama saya "Kisah". Salam kenal. Mari belajar "Golang" bersama HMSI.`
	fmt.Println(pesan)

	// Zero Value & Nil
	var address string = ""
	var isBootcamp bool
	fmt.Printf("Address %s \n", address)
	fmt.Printf("IsBootcamp %t \n", isBootcamp)

	// ========================================================================================================
}
