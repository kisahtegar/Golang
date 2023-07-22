package main

import (
	"fmt"
	"pertemuan_7/interfaces"
	"sort"
)

func main() {
	// ================================================================================================
	// [Apa Itu Interface ?]
	// Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang
	// dibungkus dengan nama tertentu. Interface merupakan tipe data. Nilai objek bertipe interface
	// zero value-nya adalah nil. Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek
	// konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

	// ------------------------------------------------------------------------------------------------
	println("======== Penerapan Interface ========")
	// [1] Penerapan Interface
	//
	// Berikut beberapa contoh penggunaan interface

	// -------[Contoh : 1] -------

	var infoCustomer interfaces.Information // interface
	var dataCustomer = interfaces.Customer{
		Id:      1,
		Name:    "Budi",
		Address: "Tangerang",
		Contact: interfaces.Contact{
			Phone: "08214491525",
			Email: "budi@mail.com",
		},
	}
	infoCustomer = &dataCustomer

	customerName := infoCustomer.GetName()
	customerContacts := infoCustomer.GetContact()

	fmt.Println("Customer name is ", customerName)
	for key, v := range customerContacts {
		fmt.Printf("Customer %s is %s\n", key, v)
	}
	fmt.Println()

	// ------- [Contoh : 2] -------

	var dataCustomer2 = interfaces.Customer{
		Id:      2,
		Name:    "Farhan",
		Address: "Jakarta",
	}
	CustomerInfo := interfaces.NewCustomer(&dataCustomer2) // include interface
	customerName2 := CustomerInfo.GetName()

	fmt.Println("Customer name is", customerName2)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	println("========== Interface Kosong ==========")
	// [2] Interface Kosong & type alias "any"
	//
	// Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any, merupakan
	// tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan
	// array, pointer, dan apapun.
	// interface{} merupakan tipe data, sehingga cara penggunaannya sama seperti pada tipe data lainnya,
	// hanya saja nilai yang diisikan bisa apa saja.

	// ------- [Contoh : 1] -------
	// var interfaces2 interface{}
	var interfaces2 any // contoh dari type alias untuk interface

	interfaces2 = "Burhan"
	fmt.Println(interfaces2)
	interfaces2 = 12345
	fmt.Println(interfaces2)
	interfaces2 = true
	fmt.Println(interfaces2)
	fmt.Println()

	// ------- [Contoh : 2] -------
	var data = map[string]interface{}{
		"name":    "Kisah",
		"age":     10,
		"Hobbies": []string{"coding", "makan"},
	}
	fmt.Println(data)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	println("== Casting variable Interface Kosong ==")
	// [3] Casting variable Interface Kosong
	//
	// Variabel bertipe interface{} bisa ditampilkan ke layar sebagai string dengan memanfaatkan fungsi
	// print, seperti fmt.Println(). Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah
	// nilai asli, melainkan bentuk string dari nilai aslinya.
	// Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan nilai asli pada variabel
	// yang bertipe interface{}, diperlukan casting ke tipe aslinya.

	// ------- [Contoh 1: Casting interface kosong ke int] -------
	var mineInterface any = 10
	// mineInterface = 10
	var hasilPerkalian = mineInterface.(int) * 2
	fmt.Println(mineInterface, "*", 2, "=", hasilPerkalian)

	// ------- [Contoh 2: Casting interface kosong ke objek Pointer] -------
	var mineInterface2 any = &interfaces.Customer{Id: 4, Name: "Anton"}
	var name = mineInterface2.(*interfaces.Customer).Name
	fmt.Println(name)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	println("=== Implementasi Kontrak Interface ===")
	// [3] Implementasi Kontrak Interface
	//
	// Package sort adalah package yang berisikan utilitas untuk proses pengurutan
	// Agar data bisa diurutkan, kita harus implementasikan kontrak interface pada sort.Interface
	// Visit https://pkg.go.dev/sort

	customers := []interfaces.Customer{
		{
			Id:   1,
			Name: "Kisah",
		},
		{
			Id:   2,
			Name: "Faiz",
		},
		{
			Id:   3,
			Name: "Aladdin",
		},
		{
			Id:   4,
			Name: "Xurma",
		},
	}

	fmt.Println(customers)
	sort.Sort(interfaces.CustomerSlice(customers))
	fmt.Println(customers)
	fmt.Println()

	// ================================================================================================
	println("=============== Study Cases ===============")
	// Umur Ayu adalah 10 tahun
	ayu := interfaces.Anggota{Nama: "Ayu", Umur: 10}

	fmt.Println("a. Berapa umur cinta, jika ayu berumur 10 tahun?")
	// karena umur Ayu 5 tahun dibawah cinta dan ayu berumur 10 tahun.
	// Jadi umur cinta, 10 + 5 = 15 Tahun
	cinta := interfaces.Anggota{
		Nama: "Cinta",
		Umur: ayu.GetUmur() + 5,
	}
	fmt.Printf("   - Umur %s adalah %d tahun\n\n", cinta.Nama, cinta.Umur)

	fmt.Println("b. Berapa umur budi dan ibu,jika ayu berumur 10 tahun?")
	// - Budi 2 tahun diatas Cinta, Umur cinta 15 tahun maka umur budi,
	//   15+2 = 17 Tahun,
	budi := interfaces.Anggota{
		Nama: "Budi",
		Umur: cinta.GetUmur() + 2,
	}
	fmt.Printf("   - Umur %s adalah %d tahun\n", budi.Nama, budi.Umur)

	// - Umur ibu 3 kali lebih besar dari ayu dan ditambah 2. 3*10+2 = 32
	ibu := interfaces.Anggota{
		Nama: "Ibu",
		Umur: 3*ayu.GetUmur() + 2,
	}
	fmt.Printf("   - Umur %s adalah %d tahun\n\n", ibu.Nama, ibu.Umur)

	fmt.Println("c. Berapa umur ayah,jika ayu berumur 10 tahun?")
	// menambahkan keseluruhan umur anak
	ayah := interfaces.Anggota{
		Nama: "Ayah",
		Umur: ayu.GetUmur() + budi.GetUmur() + cinta.GetUmur(),
	}
	fmt.Printf("   - Umur %s adalah %d tahun\n\n", ayah.Nama, ayah.Umur)

	fmt.Println("d. Berdasarkan perhitungan diatas, urutkan berdasarkan anak tertua hingga termuda?")
	totalKeluarga := interfaces.SortKeluarga{
		ayu,
		budi,
		cinta,
	}
	sort.Sort(totalKeluarga) // sorting keluarga
	for _, v := range totalKeluarga {
		fmt.Printf("   - Umur %s adalah %d tahun\n", v.GetNama(), v.GetUmur())
	}
}
