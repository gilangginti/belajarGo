package main

import "fmt"

func getGoodBy(name string) string {
	return "Goodby " + name
}
func main() {
	by := getGoodBy
	fmt.Println(by("Gilang"))
}
