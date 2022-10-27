package main

import "fmt"

func Assortions() interface{} {
	return true
}

func main() {
	var result interface{} = Assortions()
	// var convertRes string = result.(string)
	// fmt.Println(convertRes)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is String")
	case int:
		fmt.Println("Value", value, "Is Integer")
	default:
		fmt.Println("Unknow Type")
	}
}
