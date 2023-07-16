package main

import (
	"fmt"
)

func main() {
	// ================================================================================================
	println("=============== Pointer ==============")
	// [Apa Itu Pointer]
	// Pointer adalah reference atau alamat memori.
	// Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu
	// sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada
	// variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.
	// Secara default GO menerapkan “pass by value”, untuk menggunakan “pass by reference” maka GO
	// menyediakan Pointer.
	//
	// Source: https://www.scaler.com/topics/difference-between-call-by-value-and-call-by-reference/

	// ------------------------------------------------------------------------------------------------
	// [1] Penerapan Pointer
	//
	// Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe
	// data ketika deklarasi.
	// Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai
	// yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.
	// Ada dua hal penting yang perlu diketahui mengenai pointer:
	// 	+ Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda [ampersand (&)]
	// 	  tepat sebelum nama variabel. Metode ini disebut dengan [referencing].
	// 	+ Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda
	// 	  [asterisk (*)] tepat sebelum nama variabel. Metode ini disebut dengan [dereferencing].

	var angkaPertama int = 4
	var angkaKedua *int = &angkaPertama

	fmt.Println("Angka pertama (value) : ", angkaPertama)
	fmt.Println("Angka pertama (address) : ", &angkaPertama)
	fmt.Println("Angka kedua (value) : ", *angkaKedua)
	fmt.Println("Angka kedua (address) : ", angkaKedua)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [2] Parameter Pointer
	//
	// Pointer bisa juga untuk dijadikan sebagai parameter pada sebuah fungsi, caranya sama.
	// Catatan : pointer adalah pass by reference, maka jika variable diubah, maka variable yang
	// mempunya alamat memori yang sama akan mempunyai nilai yang sama juga.
	var angka int = 4
	fmt.Println("Angka default :", angka)
	ubah(&angka, 6)
	fmt.Println("Angka setelah diubah dengan func (Past by reference) :", angka)
	fmt.Println()

	// ================================================================================================
	println("================ Struct ==============")
	// [Apa Itu Struct]
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi , yang dibungkus sebagai
	// tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi.
	// Bahkan, kita bisa membuat struct di dalam struct.
	//
	// Konsep struct di golang mirip dengan konsep class pada OOP. Dengan memanfaatkan struct, grouping
	// data akan lebih mudah, selain itu rapih dan gampang untuk di-maintain.

	// ------------------------------------------------------------------------------------------------
	// [1] Deklarasi & Inisialisai Struct

	// Untuk deklarasi struct, silahkan gunakan keyword type, seperti contoh dibawah
	// var cust customer
	var cust = customer{} // initialize struct

	cust.Id = 1
	cust.Name = "Kisah"
	cust.Address = "Tangerang"

	// Cara inisialisasi variabel objek adalah dengan menambahkan kurung kurawal setelah nama struct.
	// Nilai masing-masing property bisa diisi pada saat inisialisasi. var cust = customer{}
	var cust2 = customer{Id: 2, Name: "Tono"}
	fmt.Println(cust2)

	// ------------------------------------------------------------------------------------------------
	// [2] Embedded Struct
	//
	// Struct contact, diembed ke struct customer

	cust.contact.Phone = "0821421414"
	cust.contact.Email = "kisah@unipi.com"

	fmt.Println(cust)
	fmt.Println(cust.Name)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	// [3] Kombinasi Slice & Struct
	//
	// Slice dan struct bisa dikombinasikan, caranya adalah dengan menggunakan tanda [] sebelum tipe
	// data pada saat deklarasi.

	var customers = []customer{
		{
			Id:      1,
			Name:    "Budi",
			Address: "Tangerang",
			contact: contact{
				Phone: "08211319734",
				Email: "budi@mail.com",
			},
		},
		{
			Id:      2,
			Name:    "Munar",
			Address: "Jakarta",
			contact: contact{
				Phone: "08284519734",
				Email: "munar@mail.com",
			},
		},
	}

	for _, v := range customers {
		fmt.Printf("Nama customer %s, Alamat %s, kontak: phone(%s), email(%s)\n", v.Name, v.Address, v.Phone, v.Email)
	}

	// ================================================================================================
	println("\n================ Method ==============")
	// [Apa Itu Method]
	// Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa
	// diakses lewat variabel objek.
	// Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level
	// private.

	// ------------------------------------------------------------------------------------------------
	// [1] Penerapan Method
	// Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi, ditentukan
	// juga siapa pemilik method tersebut.

	var customer = customer{
		Id:      5,
		Name:    "Fullan",
		Address: "Banten",
		contact: contact{
			Phone: "08288425734",
			Email: "fullan@mail.com",
		},
	}

	customersName := customer.getName()
	customersName2 := customer.getNamePastByReference()
	customersContact := customer.getContact()

	fmt.Println("Customer name is", customersName)
	fmt.Println("Customer name is(past  by reference)", customersName2)

	for key, val := range customersContact {
		fmt.Printf("Customer %s is %s\n", key, val)
	}

	// ================================================================================================
	// println("=============== Study Cases ===============")
}

// This function refers to ([2] Parameter Pointer)
func ubah(number *int, nilai int) {
	*number = nilai
}

// This customer struct type refers to (STRUCT)
type customer struct {
	Id      uint16
	Name    string
	Address string
	contact
}

// This contact struct type refers to (STRUCT)
type contact struct {
	Phone, Email string
}

// This method is used to get name from customer struct.
func (cust customer) getName() string {
	return cust.Name
}

func (cust *customer) getNamePastByReference() string {
	cust.Name = "change name"
	return cust.Name
}

// This method is used to get contact from customer embed struct contact.
func (cust customer) getContact() map[string]string {
	return map[string]string{
		"Phone": cust.contact.Phone,
		"Email": cust.contact.Email,
	}
}
