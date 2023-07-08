package studyCases

import "fmt"

func Aritmatika1() {
	fmt.Println("\n[1] Study case: Soal test bootcamp")
	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 1)
	fmt.Println("\n===== [Soal No 1] =====")
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

// Study case 3 july
func Aritmatika2() {
	fmt.Println("\n[2] Study case: Soal 3 july")
	// ------------------------------------------------------------------------------------------------
	// (Soal Nomor 1)
	fmt.Println("\n========== [Soal No 1] ==========")
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
