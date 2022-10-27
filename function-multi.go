package main

import "fmt"

func myFunc(name string, age int) (string, int) {
	return name, age
}

func main() {
	// variabel kedua jika tidak digunakan pakai _
	name, age := myFunc("gilang", 100)
	fmt.Println("nama: ", name, "age: ", age)
}
