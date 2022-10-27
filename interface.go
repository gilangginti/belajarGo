package main

import "fmt"

type HasName interface {
	getName() string
}

func sayMe(hasName HasName) {
	fmt.Println("name", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}
func main() {
	var data Person
	data.Name = "Gilang"

	sayMe(data)
}
