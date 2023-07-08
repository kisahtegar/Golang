package main

import (
	"fmt"
	studyCases "pertemuan_3/study_cases"
)

func main() {
	// ================================================================================================
	// [OPERATOR ARITMATIKA]
	// Operator yang digunakan untuk operasi yang sifatnya perhitungan.
	//
	// 		Tanda	Keterangan
	// 		+ 		penjumlahan
	// 		-		pengurangan
	// 		*		perkalian
	// 		/ 		pembagians
	// 		%		Modulus / sisa hasil pembagian
	//

	var nilai = (((2+6)%3)*4 - 2) / 3
	fmt.Println("Nilai dari (((2 + 6) % 3) * 4 - 2) / 3 = ", nilai)

	// ================================================================================================
	// [OPERATOR PERBANDINGAN]
	// Operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi. Hasilnya berupa nilai
	// boolean, true atau false
	//
	// 		Tanda	Keterangan
	// 		==		Apakah nilai kiri sama dengan nilai kanan
	// 		!=		Apakah nilai kiri tidak sama dengan nilai kanan
	// 		<		Apakah nilai kiri lebih kecil daripada nilai kanan
	// 		<=		Apakah nilai kiri lebih kecil atau dama dengan nilai kanan
	// 		>		Apakah nilai kiri lebih besar daripada nilai kanan
	// 		>=		Apakah nilai kiri lebih besar atau dama dengan nilai kanan
	//

	// ================================================================================================
	// [OPERATOR LOGIKA]
	// Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool  (bisa berupa
	// variabel bertipe bool, atau hasil dari operator perbandingan).
	//
	// 		Tanda	Keterangan
	// 		&&		Kiri dan kanan
	// 		||		Kiri atau kanan
	// 		!		Nilai kebalikan (negasi)
	//

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("Left && right :\t(%t)\n", leftAndRight)

	result := !left
	fmt.Println("Result : ", result)

	// ================================================================================================
	// [OPERATOR ASSIGNMENT]
	//
	// 		Tanda	Keterangan
	// 		=		Menetapkan nilai dari operan sisi kanan ke operan sisi kiri
	// 		+=		menambahkan operan kanan ke operan kiri dan menetapkan hasilnya ke operan kiri
	// 		-=		mengurangkan operan kanan ke operan kiri dan menetapkan hasilnya ke operan kiri
	// 		*=		mengalikan operan kanan ke operan kiri dan menetapkan hasilnya ke operan kiri
	// 		/=		membagikan operan kanan ke operan kiri dan menetapkan hasilnya ke operan kiri
	//

	var tinggiBadan uint8 = 179
	tinggiBadan += 11
	fmt.Printf("Tinggi Badan : %d \n", tinggiBadan)

	// ================================================================================================
	// [Study case: Soal test bootcamp.]
	//
	fmt.Println("\n==================[Study case]==================")
	studyCases.Aritmatika1()
	fmt.Println("\n------------------------------------------------")
	studyCases.Aritmatika2()

}
