package main

import (
	"fmt"
)

func main() {
	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 1)
	fmt.Println("========== [Soal No 1] ==========")
	var uangSiti int = 5000
	var kembalian int = 1500

	hargaCokelat := uangSiti - kembalian

	fmt.Printf("Harga cokelat = %d - %d = %d\n\n", uangSiti, kembalian, hargaCokelat)

	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 2)
	fmt.Println("========== [Soal No 2] ==========")
	var hargaPensil int = 2000
	var pensilDibeli int = 2
	var uangAnton int = 10000

	totalHargaPensil := hargaPensil * pensilDibeli
	uangKembalianAnton := uangAnton - totalHargaPensil

	fmt.Printf("Uang Kembalian Anton = %d - %d = %d\n\n", uangAnton, totalHargaPensil, uangKembalianAnton)

	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 3)
	fmt.Println("========== [Soal No 3] ==========")
	var hargaBuku int = 3500
	var uangNana int = 5000

	uangKembalianNana := uangNana - hargaBuku

	fmt.Printf("Uang Kembalian Nana = %d - %d = %d\n\n", uangNana, hargaBuku, uangKembalianNana)

	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 4)
	fmt.Println("========== [Soal No 4] ==========")
	var uangDedi int = 20000
	var hargaEskrim int = 7500
	var hargaSusu int = 5000

	totalHargaDedi := hargaEskrim + hargaSusu
	uangKembalianDedi := uangDedi - totalHargaDedi

	fmt.Printf("Uang Kembalian Dedi : %d - %d = %d\n\n", uangDedi, totalHargaDedi, uangKembalianDedi)

	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 5)
	fmt.Println("========== [Soal No 5] ==========")
	var hargaMangga int = 35000
	var uangMonica int = 50000

	kembalianMonica := uangMonica - hargaMangga

	fmt.Printf("Uang Kembalian Monica : %d - %d = %d\n\n", uangMonica, hargaMangga, kembalianMonica)
}
