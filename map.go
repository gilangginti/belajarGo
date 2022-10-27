package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "gilang",
		"age":  "20",
	}
	person["address"] = "jl. Lorem ipsum"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	var books map[string]string = make(map[string]string)
	books["id"] = "1"
	books["title"] = "Buku A"

	fmt.Println(books)
	fmt.Println(books["id"])
	fmt.Println(books["title"])

	delete(books, "title")
	fmt.Println(books)
}
