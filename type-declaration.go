package main

import "fmt"

func main() {
	//noHP adalah alias untuk tipe data string
	type NoHP string

	
	//buat variable dengan tipe noHP yang alias dari string
	var no_alung NoHP = "+61212121212"
	fmt.Println(no_alung)
}