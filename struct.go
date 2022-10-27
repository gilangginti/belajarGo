package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var data Customer
	data.Name = "Gilang"
	data.Address = "jl. Lorem"
	data.Age = 10
	data.sayHi("wakwaw")
	data2 := Customer{
		Name:    "Wk",
		Address: "jl. l",
		Age:     20,
	}

	fmt.Println("Data", data)
	fmt.Println("Data", data2)
}
