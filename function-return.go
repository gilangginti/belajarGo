package main

import "fmt"

func myFunc(name string) string {
	return "nama:" + name
}

func main() {
	name := "gilang"
	result := myFunc(name)

	fmt.Println(result)
}
