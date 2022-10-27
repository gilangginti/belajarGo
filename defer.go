package main

import "fmt"

func logging() {
	fmt.Println("Logging true")
}

func runApp() {
	defer logging()
	fmt.Println("App berhasil dijalankan")
}
func main() {
	runApp()
}
