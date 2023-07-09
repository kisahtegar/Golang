package main

import (
	"fmt"
	studyCase "pertemuan_4/study_case"
)

func main() {
	// ================================================================================================
	println("================== ARRAY ==================")
	// [ARRAY]
	// Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki
	// kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan
	// di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan. Default nilai tiap
	// elemen array pada awalnya tergantung dari tipe datanya. Jika int maka tiap element zero value-nya
	// adalah 0, jika bool maka false, dan seterusnya. Setiap elemen array memiliki indeks berupa angka
	// yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0.

	// ------------------------------------------------------------------------------------------------
	// [1] Contoh penerapan array.

	var names1 [3]string
	names1[0] = "Kisah"
	names1[1] = "Tegar"
	names1[2] = "Putra"
	// names1[3] = "test" // Kode ini akan menghasilkan error, melebihi alokasi awal

	fmt.Println(names1[0], names1[1], names1[2])
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Inisialisasi nilai awal Array.
	//
	// Pengisian elemen array bisa dilakukan pada saat deklarasi variabel. Caranya dengan menuliskan data
	// elemen dalam kurung kurawal setelah tipe data, dengan pembatas antar elemen adalah tanda koma (,).
	// Penggunaan fungsi fmt.Println() pada data array tanpa mengakses indeks tertentu, akan menghasilkan
	// output dalam bentuk string dari semua array yang ada. Teknik ini biasa digunakan untuk debugging
	// data array.

	var names2 = [3]string{"Kisah", "Tegar", "Putra"}

	fmt.Println("Jumlah element:\t\t", len(names2))
	fmt.Println("Isi semua element:\t", names2)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [3] Looping array menggunakan keyword (for).
	//
	// Keyword for dan array memiliki hubungan yang sangat erat. Dengan memanfaatkan perulangan menggunakan
	// keyword ini, elemen-elemen dalam array bisa didapat. Ada beberapa cara yang bisa digunakan untuk
	// perulangan data array, cara diatas adalah dengan memanfaatkan variabel iterasi perulangan untuk
	// mengakses elemen berdasarkan indeks-nya.

	var names3 = [3]string{"Kisah", "Tegar", "Putra"}

	for i := 0; i < len(names3); i++ {
		fmt.Printf("Elemen %d: %s\n", i, names3[i])
	}
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [4] Looping array menggunakan keyword (for - range).
	//
	// Array names diambil elemen-nya secara berurutan. Nilai tiap elemen ditampung variabel oleh name
	// (tanpa huruf s), sedangkan indeks nya ditampung variabel i.
	// Output code ini, sama dengan output code sebelumnya, hanya cara yang digunakan berbeda.

	var names4 = [3]string{"Kisah", "Tegar", "Putra"}

	for i, name := range names4 {
		fmt.Printf("Element %d: %s\n", i, name)
	}
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [5] Keyword make.
	//
	// Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make.

	var names5 = make([]string, 3)
	names5[0] = "Kisah"
	names5[1] = "Tegar"
	names5[2] = "Putra"

	fmt.Println(names5)
	fmt.Println()

	// ================================================================================================
	println("================== SLICE ==================")
	// -- [SLICE] --
	// Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi
	// sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di
	// tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

	// ------------------------------------------------------------------------------------------------
	// [1] Inisialisasi slice.
	//
	// Cara pembuatan slice mirip seperti pembuatan array, bedanya tidak perlu mendefinisikan jumlah
	// elemen ketika awal deklarasi.

	var names6 = []string{"Kisah", "Tegar", "Putra"}
	fmt.Println(names6)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Fungsi append().
	//
	// Fungsi append() digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan
	// setelah indeks paling akhir.

	var names7 = []string{"Kisah", "Tegar", "Putra"}
	var nameSecond = append(names7, "S.kom")

	fmt.Println(names7)
	fmt.Println(names7[1])
	fmt.Println(nameSecond)
	fmt.Println()

	// ================================================================================================
	println("=================== MAP ===================")
	// -- [MAP] --
	// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value)
	// yang disimpan, disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier)
	// untuk pengaksesan value yang bersangkutan.
	// Kalau dilihat, map mirip seperti slice, hanya saja indeks yang digunakan untuk pengaksesan bisa
	// ditentukan sendiri tipe-nya (indeks tersebut adalah key).

	// ------------------------------------------------------------------------------------------------
	// [1] Pengunaan Map.
	//
	// Cara menggunakan map cukup dengan menuliskan keyword map diikuti tipe data key dan value-nya.
	// Zero value dari map adalah nil, maka tiap variabel bertipe map harus di-inisialisasi secara explisit
	// nilai awalnya (agar tidak nil).

	// var months map[int]string
	// months = map[int]string{}

	var months = map[string]string{}

	months["Jan"] = "January"
	months["Feb"] = "February"

	fmt.Println(months)        // map[Feb:February Jan:January]
	fmt.Println(months["Feb"]) // February
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Inisialisai nilai Map.
	//
	// Nilai variabel bertipe map bisa didefinisikan di awal, caranya dengan menambahkan kurung kurawal
	// setelah tipe data, lalu menuliskan key dan value di dalamnya. Cara ini sekilas mirip dengan definisi
	// nilai array/slice namun dalam bentuk key-value.

	// Cara verticaL
	var months1 = map[string]int{
		"January":  1,
		"February": 2,
	}

	fmt.Println("Month", months1)
	fmt.Println()

	// Variabel map bisa di-inisialisasi dengan tanpa nilai awal, caranya menggunakan tanda kurung kurawal,
	// contoh: map[string]int{}. Atau bisa juga dengan menggunakan keyword make dan new.

	// var months3 = map[string]int{}
	// var months4 = make(map[string]int)
	// var months5 = *new(map[string]int)

	// ------------------------------------------------------------------------------------------------
	// [2] Iterasi item Map menggunakan (for - range)
	//
	// Item variabel map bisa di iterasi menggunakan for - range. Cara penerapannya masih sama seperti
	// pada slice, pembedanya data yang dikembalikan di tiap perulangan adalah key dan value, bukan
	// indeks dan elemen.

	var months6 = map[string]int{"January": 1, "February": 2, "March": 3}

	for key, value := range months6 {
		fmt.Println(key, " \t:", value)
	}
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Menghapus item Map.
	//
	// Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada variabel map.

	var months7 = map[string]int{"January": 1, "February": 2, "March": 3}
	months7["April"] = 4

	delete(months7, "February")
	fmt.Println(months7)
	fmt.Println()

	// ================================================================================================
	println("=============== Study Cases ===============")
	// -- [STUDY CASES] --
	studyCase.Main()
}
