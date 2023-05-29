package main

import "fmt"

func main() {
	//string pertama tipe data untuk key, yang kedua tipe data untuk value
	person := map[string]string{
		"name":    "Lulung",
		"address": "Jember",
	}
	
	//kalau membuat variabel menggunakan var bisa namun akan menjadi lebih panjang contohnya seperti di bawah ini
	//var person1 map[string]string = map[string]string{
	//	"name":    "Lulung",
	//	"address": "Jember",
	//}

	//bisa juga menambah atau merubah datanya seperti di bawah ini
	//Ingat rumusnya kalau nambah itu keynya harus beda, kalau ubah keynya harus sama
	person["title"] = "programmer"
	person["name"] = "Alung"

	//fmt.Println(person)
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(person["address"])
	//fmt.Println(person)


	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "eko"
	book["ups"] = "YNKTS"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	
	fmt.Println(len(book))
	fmt.Println(book)
}