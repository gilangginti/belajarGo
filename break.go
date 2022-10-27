package main

import "fmt"

func main() {
	number := 1

	for number <= 30 {
		fmt.Println("hitung", number)
		if number == 28 {
			break
		}
		number++
	}
}
