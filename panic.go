package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}
