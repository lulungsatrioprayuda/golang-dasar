package main

import "fmt"

func main() {
	//variable yang isi bawaannya gak bisa di ubah
	const firstName string = "Lulung"
	const lastName = "lulung"

	const (
		alung string = "nama pertama"
		satrio = "nama tengah"
	)

	fmt.Println(satrio)
}