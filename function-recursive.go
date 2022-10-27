package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRec(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	res := factorialLoop(5)
	fmt.Println(res)
	fmt.Println(factorialRec(5))
}
