package studyCase

import "fmt"

func TriangleNumber() {
	// Buatlah sususan angka berbentuk segitiga seperti pada gambar !!
	// Tips: Gunakan perulangan bersarang, for didalam for

	fmt.Println("\n[1] STUDY CASE 1: FIZZBUZZ")

	var nomor int = 10

	for i := 1; i <= nomor; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}

func KelipatanPerulangan() {
	fmt.Println("\n[2] STUDY CASE 2: Kelipatan perulangan")
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			// multiple of both 3 and 5
			fmt.Println("Insan Pembangunan")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Insan")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Pembangunan")
			continue
		}
		fmt.Println(i)
	}
}

func PerhitunganPecahan() {
	fmt.Println("\n[3] STUDY CASE 3: Perhitungan pecahan")

	var totalBelanja int = 13500
	var uangDiberikan int = 100000

	fmt.Println("Total Belanja:", totalBelanja)
	fmt.Println("Uang Diberikan:", uangDiberikan)

	kembalian := uangDiberikan - totalBelanja

	fmt.Println("Uang kembalian yang diterima:", kembalian)

	pecahan50K := kembalian / 50000
	kembalian %= 50000

	pecahan20K := kembalian / 20000
	kembalian %= 20000

	pecahan10K := kembalian / 10000
	kembalian %= 10000

	pecahan5K := kembalian / 5000
	kembalian %= 5000

	pecahan2K := kembalian / 2000
	kembalian %= 2000

	pecahan1K := kembalian / 1000
	kembalian %= 1000

	pecahan500 := kembalian / 500
	kembalian %= 500

	pecahan200 := kembalian / 200
	kembalian %= 200

	pecahan100 := kembalian / 100
	kembalian %= 100

	fmt.Println("\nPecahan yang diterima:")

	if pecahan50K > 0 {
		fmt.Printf("Jumlah pecahan 50.000: %d\n", pecahan50K)
	}
	if pecahan20K > 0 {
		fmt.Printf("Jumlah pecahan 20.000: %d\n", pecahan20K)
	}
	if pecahan10K > 0 {
		fmt.Printf("Jumlah pecahan 10.000: %d\n", pecahan10K)
	}
	if pecahan5K > 0 {
		fmt.Printf("Jumlah pecahan 5.000: %d\n", pecahan5K)
	}
	if pecahan2K > 0 {
		fmt.Printf("Jumlah pecahan 2.000: %d\n", pecahan2K)
	}
	if pecahan1K > 0 {
		fmt.Printf("Jumlah pecahan 1.000: %d\n", pecahan1K)
	}
	if pecahan500 > 0 {
		fmt.Printf("Jumlah pecahan 500: %d\n", pecahan500)
	}
	if pecahan200 > 0 {
		fmt.Printf("Jumlah pecahan 200: %d\n", pecahan200)
	}
	if pecahan100 > 0 {
		fmt.Printf("Jumlah pecahan 100: %d\n", pecahan100)
	}
}
