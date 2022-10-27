package main

import "fmt"

func myGu() (name string, age int) {
	name = "gilang"
	age = 100
	return
}

func main() {
	name, age := myGu()
	fmt.Println(name)
	fmt.Println(age)
}
