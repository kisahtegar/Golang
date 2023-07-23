package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ================================================================================================
	println("================ defer ===============")
	// [Apa Itu Defer]
	// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
	// Seperti yang sudah dijelaskan secara singkat di atas, bahwa defer digunakan untuk mengakhirkan
	// eksekusi baris kode dalam skope blok fungsi. Ketika eksekusi blok sudah hampir selesai, statement
	// yang di-defer dijalankan.
	// Defer bisa ditempatkan di mana saja, awal maupun akhir blok. Tetapi tidak mempengaruhi kapan waktu
	// dieksekusinya, akan selalu dieksekusi di akhir.

	// ------------------------------------------------------------------------------------------------
	// [1] Penerapan Defer

	pesan("ketoprak")
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Kombinasi Defer & IIFE
	//
	// bahwa eksekusi defer adalah di akhir blok fungsi, bukan blok lainnya seperti blok seleksi kondisi.
	// Pada contoh dibawah ini,  "terimakasih sudah pesan di warung kami" , akan tetap di eksekusi setelah
	// "Pesanan anda adalah soto" ,
	// meskipun statement defer dipergunakan dalam blok seleksi kondisi if. Hal ini karena defer eksekusinya
	// terjadi pada akhir blok fungsi (dalam contoh di atas main()), bukan pada akhir blok if.
	// Agar "terimakasih sudah pesan di warung kami"  dieksekusi dalam blok if, maka harus di bungkus dengan IIFE.

	pesanan := "soto"

	if pesanan == "soto" {
		fmt.Println("Pilihan tepat, soto kami paling sedap")
		// defer fmt.Println("Terimakasih sudah pesan di warung ini")
		func() {
			defer fmt.Println("Terimakasih sudah pesan di warung ini")
		}()
	}
	fmt.Println("Pesanan anda adalah", pesanan)
	fmt.Println()

	// ================================================================================================
	println("================ erorr ===============")
	// [Apa Itu Error]
	// error merupakan sebuah tipe. Error memiliki 1 buah property berupa method Error(), method ini
	// mengembalikan detail pesan error dalam string. Error termasuk tipe yang isinya bisa nil.

	// ------------------------------------------------------------------------------------------------
	// [1] Penerapan Error

	var (
		input         string = "ini string"
		errConvertion error
		result        int
	)

	result, errConvertion = strconv.Atoi(input)

	if errConvertion != nil {
		fmt.Println(errConvertion)
	} else {
		fmt.Println(result, "Adalah sebuah angka.")
	}
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Custom Error
	// kita bisa membuat objek error sendiri dengan menggunakan fungsi errors.New()
	// Untuk menggunakan errors.New() harus import package errors terlebih dahulu

	// avg, err := average(1, 1, 2, 4, 5, 8)
	avg, err := average()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
		fmt.Println(msg)
	}
	fmt.Println()

	// ================================================================================================
	println("================ exit ===============")
	// [Apa Itu Exit]
	// Exit digunakan untuk menghentikan program secara paksa (ingat, menghentikan program, tidak
	// seperti return yang hanya menghentikan blok kode).
	// Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah parameter bertipe numerik
	// yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status ketika program berhenti.

	defer fmt.Println("Hallo")
	os.Exit(1)
	fmt.Println("Selamat datang")

	// ================================================================================================

	// ================================================================================================
	// println("=============== Study Cases ===============")
}

// This function refers to (Penerapan defer)
func pesan(pesanan string) {
	defer fmt.Println("Terimakasih sudah pesan di warung kami")

	if pesanan == "ketoprak" {
		fmt.Println("Pilihan tepat, ketoprak kami paling enak")
		return
	}

	fmt.Println("Pesanan anda adalah", pesanan)
}

func average(numbers ...int) (float64, error) {
	var avg float64 = 0
	var total int = 0

	if len(numbers) > 0 {
		for _, number := range numbers {
			total += number
		}
		avg = float64(total) / float64(len(numbers))
	} else {
		return avg, errors.New("numbers can't be empty")
	}

	return avg, nil
}
