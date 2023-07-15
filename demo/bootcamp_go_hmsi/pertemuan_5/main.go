package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// ================================================================================================
	println("=============== Functions ==============")
	// [Apa Itu Fungsi]
	// Dalam dunia pemrograman di kenal adanya istilah DRY (kependekan dari Don’t Repeat Yourself).
	// Maksudnya adalah anda tidak perlu menuliskan banyak kode yang digunakan berulangkali, cukup
	// sekali saja sesuai kebutuhan. DRY bisa kita implementasikan dengan menggunakan Fungsi.
	// Lalu, apa itu Fungsi? Fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.

	// [Penulisan Fungsi]
	// Cara membuat fungsi cukup mudah, yaitu dengan menuliskan keyword func, diikuti setelahnya nama
	// fungsi, kurung yang berisikan parameter, dan kurung kurawal untuk membungkus blok kode.
	// Parameter sendiri adalah variabel yang disisipkan pada saat pemanggilan fungsi.

	var name = "Kisah"
	printPesertaBootcamp(name)

	// ------------------------------------------------------------------------------------------------
	// [1] Return Value
	//
	// Sebuah fungsi bisa dirancang tidak mengembalikan nilai balik (void), atau bisa mengembalikan
	// suatu nilai. Fungsi yang memiliki nilai kembalian, harus ditentukan tipe data nilai baliknya
	// pada saat deklarasi.
	//
	// Khusus untuk fungsi yang tipe data parameternya sama, bisa ditulis dengan gaya yang unik. Tipe
	// datanya dituliskan cukup sekali saja di akhir.

	firstNumber, secondNumber := 9, 8
	result := getGreater(firstNumber, secondNumber)
	fmt.Println("Nilai terbesar adalah", result)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Multiple Return
	//
	// Cara membuat fungsi yang memiliki banyak nilai balik tidaklah sulit. Tinggal tulis saja pada
	// saat deklarasi fungsi semua tipe data nilai yang dikembalikan, dan pada keyword return tulis
	// semua data yang ingin dikembalikan.

	var diameter float64 = 15
	luas, keliling := hitung(diameter)
	fmt.Printf("Luas lingkaran : %.2f\n", luas)
	fmt.Printf("Keliling lingkaran : %.2f\n\n", keliling)

	// ------------------------------------------------------------------------------------------------
	// [3] Variadic Function
	//
	// Pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud [tak terbatas] di sini adalah
	// jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.
	//
	// Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya adalah pada
	// parameter jenis ini ditambahkan tanda titik tiga kali (...)  tepat setelah penulisan variabel
	// (sebelum tipe data).

	total := getTotal(1, 2, 3, 4, 5)
	fmt.Printf("Total : %d\n\n", total)

	// Slice bisa digunakan sebagai parameter variadic. Caranya dengan menambahkan tanda titik tiga
	// kali, tepat setelah nama variabel yang dijadikan parameter

	numbersSlices := []int{1, 2, 3, 4, 5, 6, 7, 8}
	totalWithSlice := getTotal(numbersSlices...)
	fmt.Printf("Total With Slice: %d\n\n", totalWithSlice)

	// Parameter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter
	// variadic-nya harus diposisikan di akhir.

	mySports("kisah", "Football", "Racing", "Chess")
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [4] Closure Function
	//
	// Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan
	// konsep tersebut, kita bisa membuat fungsi di dalam fungsi.

	var getMinMax = func(numbers []int) (int, int) {
		var min, max int

		// Loop through the numbers
		for i, n := range numbers {
			// fmt.Println("Index : ", i)
			// fmt.Println("Number : ", n)

			if i == 0 {
				max, min = n, n
				// fmt.Println(i)
				// fmt.Println(max, min)
			} else if n > max {
				max = n
				// fmt.Println("Max :", max)
			} else if n < min {
				min = n
				// fmt.Println("Min :", min)
			}
		}

		return min, max
	}

	var numbers = []int{2, 1, 7, 4}
	var min, max = getMinMax(numbers)
	fmt.Printf("Nilai min adalah %d dan nilai max adalah %d\n", min, max)

	// (Immediately-Invoked Function Expression (IIFE))
	// Closure jenis ini dieksekusi langsung pada saat deklarasinya. Biasa digunakan untuk membungkus
	// proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak.

	var getModulus = func(number int) int {
		return number % 3
	}(11)

	fmt.Printf("Nilai sisa bagi adalah %d\n\n", getModulus)

	// ================================================================================================
	// println("=============== Study Cases ===============")

	var nama = "kasur rusak"
	hasil := isPalindrome(nama)
	if hasil {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")

	}

}

// This function refers to first example
func printPesertaBootcamp(name string) {
	fmt.Println("Nama peserta bootcamp", name)
	fmt.Println()
}

// This function refers to ([1] RETURN VALUE).
func getGreater(firstNumber, secondNumber int) int {
	var result int

	if firstNumber > secondNumber {
		result = firstNumber
	}
	if firstNumber < secondNumber {
		result = secondNumber
	}

	return result
}

// This function refers to ([2] Multiple Return)
func hitung(diameter float64) (float64, float64) {
	// hitung luas
	luas := math.Pi * math.Pow(diameter/2, 2)

	// hitung keliling
	keliling := math.Pi * diameter

	return luas, keliling
}

// This function refers to ([2] Multiple Return)
func getTotal(numbers ...int) int {
	var total int

	// Loop through numbers
	for _, number := range numbers {
		total = total + number
	}

	return total
}

// This function refers to ([2] Multiple Return)
func mySports(name string, sports ...string) {
	var sportsString string = strings.Join(sports, ",")

	fmt.Println("Name :", name)
	fmt.Println("Favorite sports :", sportsString)
}

// Buat fungsi dengan parameter string return value bool nama fungsi isPalindrome
func isPalindrome(name string) bool {
	for i := 0; i < len(name); i++ {
		// fmt.Println(i)
		// fmt.Println(string(name[i]))
		if string(name[i]) != string(name[len(name)-i-1]) {
			return false
		}
	}
	return true
}
