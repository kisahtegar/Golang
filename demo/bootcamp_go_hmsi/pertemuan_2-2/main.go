package main

import "fmt"

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
	fmt.Printf("Tinggi Badan : %d \n\n", tinggiBadan)

	// ================================================================================================
	// [Study case: Soal test bootcamp.]
	//
	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 1)
	fmt.Println("===== [Soal No 1] =====")
	var pinjaman float32 = 10000000
	var bunga float32 = 0.1
	var cicilan float32 = 12

	// Total Bunga
	totalBunga := pinjaman * bunga
	fmt.Printf("A) Total bunga : %.2f \n", totalBunga)

	// Total pinjaman dihitung dari pinjaman + total bunga
	totalPinjaman := pinjaman + totalBunga
	fmt.Printf("B) Total pinjaman :%.2f \n", totalPinjaman)

	totalCicilan := totalPinjaman / cicilan
	fmt.Printf("C) Total cicilan: %.2f \n\n", totalCicilan)

	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 2)
	fmt.Println("===== [Soal No 2] =====")
	var nilaiAbsensiDefault float32 = 18
	var nilaiAbsensiDapat float32 = 16
	nilaiAbsensiResult := nilaiAbsensiDapat / nilaiAbsensiDefault
	var bobotNilaiAbsensi float32 = 0.1

	var nilaiTugas float32 = 70
	var bobotNilaiTugas float32 = 0.2

	var nilaiUts float32 = 80
	var bobotNilaiUts float32 = 0.3

	var nilaiUas float32 = 70
	var bobotNilaiUas float32 = 0.4

	nilaiAkhir := (nilaiAbsensiResult * bobotNilaiAbsensi) + (nilaiTugas * bobotNilaiTugas) + (nilaiUts * bobotNilaiUts) + (nilaiUas * bobotNilaiUas)
	fmt.Printf("Nilai akhir Apip : %.3f \n\n", nilaiAkhir)

	// ------------------------------------------------------------------------------------------------
	// (Soal nomor 3)
	fmt.Println("===== [Soal No 3] =====")
	var jarak float32 = 63
	var habisBensin float32 = 17
	var biayaBensin float32 = 8350

	// Dit: Biaya bensin?
	totalBensin := jarak / habisBensin
	fmt.Println("Total Bensin = ", totalBensin)

	totalBiaya := biayaBensin * totalBensin
	fmt.Println("Total Biaya = ", totalBiaya)

}
