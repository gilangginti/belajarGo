package main

import "fmt"

func main() {
	person := "wakwaw"
	age := 10
	child := 2
	if person == "gilang" && age == 10 {
		if child == 1 {
			fmt.Printf("Nama dan child benar")
		} else {
			fmt.Println("nama benar , child salah")
		}

	} else if person == "wakwaw" {
		fmt.Println("Nama :", person)
	}
}
