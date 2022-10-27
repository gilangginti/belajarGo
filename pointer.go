package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddress(address *Address) {
	address.Country = "England"
}
func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"
	*address2 = Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := new(Address)

	fmt.Println("1", address1)
	fmt.Println("2", address2)
	fmt.Println("3", address3)
	fmt.Println("4", address4)

	alamat := Address{
		City:     "New York",
		Province: "Los Angles",
		Country:  "",
	}
	// var alamatPointer *Address = &alamat
	ChangeAddress(&alamat)
	fmt.Println(alamat)

}
