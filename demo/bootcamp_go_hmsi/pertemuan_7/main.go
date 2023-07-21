package main

import (
	"fmt"
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

	var infoCustomer Information // interface
	var dataCustomer = customer{
		Id:      1,
		Name:    "Budi",
		Address: "Tangerang",
		contact: contact{
			Phone: "08214491525",
			Email: "budi@mail.com",
		},
	}
	infoCustomer = dataCustomer

	customerName := infoCustomer.getName()
	customerContacts := infoCustomer.getContact()

	fmt.Println("Customer name is ", customerName)
	for key, v := range customerContacts {
		fmt.Printf("Customer %s is %s\n", key, v)
	}
	fmt.Println()

	// ------- [Contoh : 2] -------

	var dataCustomer2 = customer{
		Id:      2,
		Name:    "Farhan",
		Address: "Jakarta",
	}
	CustomerInfo := NewCustomer(&dataCustomer2) // include interface
	customerName2 := CustomerInfo.getName()

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
	// var myInterface interface{}
	var myInterface any // contoh dari type alias untuk interface

	myInterface = "Burhan"
	fmt.Println(myInterface)
	myInterface = 12345
	fmt.Println(myInterface)
	myInterface = true
	fmt.Println(myInterface)
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
	var mineInterface2 any = &customer{Id: 4, Name: "Anton"}
	var name = mineInterface2.(*customer).Name
	fmt.Println(name)
	fmt.Println()

	// ------------------------------------------------------------------------------------------------
	println("=== Implementasi Kontrak Interface ===")
	// [3] Implementasi Kontrak Interface
	//
	// Package sort adalah package yang berisikan utilitas untuk proses pengurutan
	// Agar data bisa diurutkan, kita harus implementasikan kontrak interface pada sort.Interface
	// Visit https://pkg.go.dev/sort

	customers := []customer{
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
	sort.Sort(customerSlice(customers))
	fmt.Println(customers)
	// ================================================================================================
	// println("=============== Study Cases ===============")
}

// This customer struct type.
type customer struct {
	Id      uint16
	Name    string
	Address string
	contact
}

// This contact struct type.
type contact struct {
	Phone, Email string
}

// This method is used to get name from customer struct.
func (cust customer) getName() string {
	return cust.Name
}

// This method is used to get contact from customer embed struct contact.
func (cust customer) getContact() map[string]string {
	return map[string]string{
		"Phone": cust.contact.Phone,
		"Email": cust.contact.Email,
	}
}

// this Information is interface.
type Information interface {
	getName() string
	getContact() map[string]string
}

// This method return customer information.
func NewCustomer(cust *customer) Information {
	return &customer{
		Id:      cust.Id,
		Name:    cust.Name,
		Address: cust.Address,
	}
}

// customerSlice need to implement 3 method for sort.Interface (Len, Less, Swap)
type customerSlice []customer

func (v customerSlice) Len() int {
	return len(v)
}
func (v customerSlice) Less(i, j int) bool {
	return v[i].Name < v[j].Name
}
func (v customerSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
