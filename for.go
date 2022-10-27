package main

import "fmt"

func main() {
	counter := 1

	for counter <= 20 {
		fmt.Println("Countdown", counter)
		counter++
	}

	slice := []string{"gilang", "fauzi", "aku", "adalah"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"wakwaw", "adalah", "manusia"}

	// Jika tidak digunakan maka mengguanakn _
	// Jika tidak array maka akan mereturn index
	for _, name := range names {
		fmt.Println("Index", "=", name)
	}

	// Jika make maka akan mereturn key
	person := make(map[string]string)
	person["nama"] = "Gilang"
	person["umur"] = "20"

	for key, item := range person {
		fmt.Println(key, "=", item)
	}
}
