package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, item := range numbers {
		total += item
	}
	return total
}

func main() {
	result := sumAll(100, 90, 80, 70)
	fmt.Println(result)

	slice := []int{10, 20, 30, 40}
	res := sumAll(slice...)
	fmt.Println(res)
}
