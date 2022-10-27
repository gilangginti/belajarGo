package main

import "fmt"

func Ups(i interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "String"
	}
}

func main() {
	var data interface{} = Ups("wakwaw")
	fmt.Println(data)
}
