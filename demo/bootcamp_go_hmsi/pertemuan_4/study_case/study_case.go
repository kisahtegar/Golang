package studyCase

import (
	"fmt"
	"sort"
)

func Main() {
	// ----------------------------------------------------------------
	fmt.Println("==================== SOAL 1 ====================")

	Latihan1()

	// ----------------------------------------------------------------
	fmt.Println("==================== SOAL 2 ====================")
	// Find numbers larger and show summarry and average.

	numbers := []int{11, 6, 31, 201, 99, 861, 1, 7, 14, 79}
	fmt.Println("Sederet bilangan slice :", numbers)
	value := findLargestNum(numbers)

	fmt.Printf("Bilangan terbesar : %d\n\n", value)

	// ----------------------------------------------------------------
	fmt.Println("==================== SOAL 3 ====================")
	// Urutkan bilangan-bilangan dari array berikut dengan menggunakan bubble sort

	numbersBubbleSort := [10]int{99, 2, 64, 8, 111, 33, 65, 11, 102, 50}
	fmt.Println("Sederet bilangan belum di sort :", numbersBubbleSort)
	valueBubbleSort := bubbleSortArray(numbersBubbleSort)
	fmt.Printf("Bilangan yang sudah di bubble sort : %d\n\n", valueBubbleSort)

	// ----------------------------------------------------------------
	fmt.Println("==================== SOAL 4 ====================")
	// Urutkan bilangan-bilangan dari array berikut dengan menggunakan fungsi bawaan golang

	numbersShortGolang := [10]int{55, 31, 20, 1, 10, 89, 90, 66, 11, 21}
	fmt.Println("Sederet bilangan belum di sort :", numbersShortGolang)
	sort.Ints(numbersShortGolang[:]) // Functions ini disediakan dari goolang.
	fmt.Printf("Bilangan yang sudah di sort dengan Go : %d\n\n", numbersShortGolang)
}

func Latihan1() {
	fmt.Println("\n[Study case 1: Tampilkan sumary dan average]")
	fmt.Println()

	/*
		Latihan 1:
		Ada sebuah slice numbers dibawah, Tugasnya adalah tampilkan total summary dan average

		result:
		"total summary adalah ..."
		"avarage adalah ..."
	*/

	// Variables
	numbers := []int{2, 5, 6, 8, 9, 2}
	summary := 0
	average := float64(summary) / float64(len(numbers))

	// Loop through numbers
	for _, number := range numbers {
		summary += number
	}

	// Results
	fmt.Println("Total summary adalah", summary)
	fmt.Printf("Average adalah %.1f", average)
	fmt.Println()
}

// Finding the largest number and return integer value.
func findLargestNum(numbers []int) int {
	var maxNum int
	maxNum = numbers[0] // Assign first number in slice of numbers.

	// Loop through numbers.
	for _, number := range numbers {
		// Check all number is greater than maxNum and change maxNum when number is greater.
		if number > maxNum {
			maxNum = number
		}
	}
	return maxNum
}

// Bubble sort function.
func bubbleSortArray(numbers [10]int) [10]int {
	// Loop ini akan melakukan perulangan sesuai dengan jumlah numbers.
	for i := 0; i <= len(numbers); i++ {
		// fmt.Println(i) // Uncomment untuk melihat proses perulangannya.

		// Loop ini akan mengulang sampai index number selesai
		for j := 0; j < len(numbers)-1-i; j++ {
			// fmt.Printf("numbers[j] : %d, numbers[j+1] : %d\n", numbers[j], numbers[j+1]) // Uncomment untuk melihat proses swap numbers.

			// jika number indeks pertama lebih besar dari indeks ke 2
			// maka swap terjadi
			if numbers[j] > numbers[j+1] {
				// fmt.Println(numbers) // Uncomment untuk melihat proses swap numbers.

				// syntax ini berfungsi untuk melakukan swap
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]

				// fmt.Println(numbers) // Uncomment untuk melihat proses swap numbers.
			}
		}
	}
	// jika sudah kembalikan array of numbers.
	return numbers
}
