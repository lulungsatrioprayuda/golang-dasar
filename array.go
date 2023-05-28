package main

import "fmt"

func main() {
	//buat array manual
	var names [3]string

	names[0] = "Lulung"
	names[1] = "Satrio"
	names[2] = "Prayuda"
	//data satrio bisa di manipulasi menjadi satrios
	names[1] = "Satrios"
	fmt.Println(names)
	fmt.Println(names[2])

	//buat array langsung dengan datanya
	var values = [3]int{
		90,
		10,
		11,
	}

	fmt.Println(values)
	fmt.Println(values, values[1])
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}