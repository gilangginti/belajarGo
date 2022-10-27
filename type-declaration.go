package main

import "fmt"

func main() {
	type ofName string
	type ofAge int
	var name ofName = "Muhammad"
	var age ofAge = 21
	fmt.Println(name, age)
}
