package main

import "fmt"

func main() {
	nama := getNamaLengkap("Lulung", "Satrio","Prayuda")
	umur := getUmurs(2, 10)

	fmt.Println("Nama :", nama.NamaPertama, nama.NamaTengah,nama.NamaAkhir)
	fmt.Println("Umur :", umur)
}