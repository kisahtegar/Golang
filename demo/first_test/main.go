package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	var nama string

	fmt.Println("Masukkan nama anda : ")
	fmt.Scan(&nama)

	if nama == "Eddy" {

		fmt.Println("Selamat sore", nama, "Selamat datang di golang.")
	} else {
		fmt.Println("Maaf nama tidak terdaftar")
	}

}
