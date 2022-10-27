package main

import "fmt"

func myFunc(name string, age int) {
	fmt.Println(name, age)
}

func main() {
	name := "gilang"
	age := 21
	myFunc(name, age)
}
