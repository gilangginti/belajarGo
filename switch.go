package main

import "fmt"

func main() {
	name := "gilang"

	// switch name {
	// case "gilang":
	// 	fmt.Println("nama nya benar")
	// case "wakwaw":
	// 	fmt.Println("aisiteru")
	// default:
	// 	fmt.Println("ini default")
	// }

	switch {
	case name == "gilang":
		fmt.Println("nama benar gilang")
	case name == "wakwaw":
		fmt.Println("nama salah")
	}
}
