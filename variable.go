package main

import "fmt"

func main() {
	var name string

	//menggunakan inisialisai tipe data
	name = "Alung"
	fmt.Println(name)
	name = "Alunga"
	fmt.Println(name)

	//langusn mengisi data tanpa inisialisasi
	var nama = "Alungsaa"
	fmt.Println(nama)

	//variable tanpa menggunakan var
	alias := "Alungs"
	fmt.Println(alias)

	//banyak variable tapi pakai 1 var
	var (
		firstName = "Alung"
		lastName = "Satrio"
	)
	
	fmt.Println(firstName , lastName)
}