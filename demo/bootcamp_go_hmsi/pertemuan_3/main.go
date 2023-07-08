package main

import (
	"fmt"
	studyCase "pertemuan_3/study_cases"
)

func main() {
	// ================================================================================================
	// [IF ELSE STATEMENT]
	// Cara penerapan if-else di GO, sama seperti dibahasa pemrograman lainnya. Yang membedakan adalah
	// tanda kurung tidak perlu dituliskan
	//
	// Note :
	// 	- GO tidak mendukung seleksi menggunakan ternary
	// 	- Kurung kurawal harus dituliskan meskipun isinya hanya satu baris kode
	fmt.Println("[IF ELSE STATEMENT]")

	var pointA int = 6

	if pointA == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if pointA >= 6 {
		fmt.Println("Lulus")
	} else if pointA == 5 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus nilai anda %d\n", pointA)
	}

	// ================================================================================================
	// [VARIABLE TEMPORARY PADA IF-ELSE]
	// Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi di mana
	// ia ditempatkan saja.
	//
	// Note :
	//	- Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference
	// 	- Variable temporary hanya bisa digunakan pada blok seleksi kondisi itu saja
	fmt.Println("\n[VARIABLE TEMPORARY PADA IF-ELSE]")

	var pointB float32 = 7850.0

	if percent := pointB / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// ================================================================================================
	// [SWITCH CASE]
	// Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel, lalu kemudian di-cek
	// nilainya.

	// Note :
	// Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya,
	// meskipun tidak ada keyword break. Konsep ini berkebalikan dengan switch pada umumnya, yang
	// ketika sebuah case terpenuhi, maka akan tetap dilanjut mengecek case selanjutnya kecuali ada
	// keyword break.
	// Jika memang ingin memaksa melanjutkan ke case selanjutnya, silahkan gunakan keyword fallthrough.

	fmt.Println("\n[SWITCH CASE]")

	var pointC = 8

	switch pointC {
	case 8:
		fmt.Println("Perfect")
		fallthrough // force to get next cases.
	case 7:
		fmt.Println("Awesome")
		fallthrough
	default:
		fmt.Println("Not bad")
	}

	// ================================================================================================
	// [SELEKSI KONDISI BERSARANG]
	// Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi.

	fmt.Println("\n[SELEKSI KONDISI BERSARANG]")

	var pointD = 0

	if pointD >= 7 {
		switch pointD {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if pointD >= 5 {
			fmt.Println("Not bad")
		} else if pointD == 4 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if pointD == 0 {
				fmt.Println("Try harder!")
			}
		}
	}

	// ================================================================================================
	// [LOOPING]
	// Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti, selama kondisi yang
	// dijadikan acuan terpenuhi.
	// Di GO keyword perulangan hanya for saja. Namun kemampuannya merupakan gabungan dari for, while
	// dan foreach pada pemrograman lain.
	//
	// Note: Keyword for - range, akan dibahas pada pertemuan berikutnya pada pembahasan array, slice dan map

	fmt.Println("\n[LOOPING]")

	for i := 0; i < 5; i++ {
		fmt.Println("Urutan nomor ke-", i)
	}

	fmt.Println("-------------------")
	var i int = 5

	for i < 10 {
		fmt.Println("Urutan nomor ke-", i)
		i++
	}

	// ================================================================================================
	// [BREAK & CONTINUE]
	// Keyword break digunakan untuk menghentikan secara paksa sebuah perulangan, sedangkan continue
	// dipakai untuk memaksa maju ke perulangan berikutnya.

	fmt.Println("\n[BREAK & CONTINUE]")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // lanjut
		}

		if i > 8 {
			break // berhenti
		}
		fmt.Println("Angka Ganjil adalah ", i)
	}

	// ================================================================================================
	// [STUDY CASE]
	fmt.Println("\n==================[Study case]==================")
	studyCase.TriangleNumber()
	fmt.Println("\n------------------------------------------------")
	studyCase.KelipatanPerulangan()
	fmt.Println("\n------------------------------------------------")
	studyCase.PerhitunganPecahan()

}
