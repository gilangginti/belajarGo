package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "MR ." + man.Name
}

func main() {
	var data Man
	data.Name = "Gilang"

	datas := Man{"Wakwaw"}
	datas.Married()
	data.Married()

	fmt.Println(data.Name)
	fmt.Println(datas.Name)
}
