package interfaces

// This customer struct type.
type Customer struct {
	Id      uint
	Name    string
	Address string
	Contact
}

// This contact struct type.
type Contact struct {
	Phone, Email string
}

// This method is used to get name from customer struct.
func (cust *Customer) GetName() string { return cust.Name }

// This method is used to get contact from customer embed struct contact.
func (cust *Customer) GetContact() map[string]string {
	return map[string]string{
		"Phone": cust.Contact.Phone,
		"Email": cust.Contact.Email,
	}
}

func (cust *Customer) GetAddress() string { return cust.Address }

// func (cust *Customer) GetId() uint {
// 	return cust.Id
// }

// this Information is interface.
type Information interface {
	GetName() string
	GetContact() map[string]string
	GetAddress() string
}

// This method return customer information.
func NewCustomer(cust *Customer) Information {
	return &Customer{
		Id:      cust.Id,
		Name:    cust.Name,
		Address: cust.Address,
	}
}

// Interface untuk Family
type Keluarga interface {
	GetNama() string
	GetUmur() uint
}

// Struct untuk anggota keluarga
type Anggota struct {
	Nama string
	Umur uint
}

// function untuk mengambil data nama anggota
func (a Anggota) GetNama() string { return a.Nama }

// function untuk mengambil data umur anggota
func (a Anggota) GetUmur() uint { return a.Umur }
