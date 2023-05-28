package main

import "fmt"

func main() {
	var name1 = "alung"
	var name2 = "Alung"

	//kalau comparation string itu sensitive case pengaruh beda huruf kecil atau besar maka hasilnya false
	var result = name1 == name2
	fmt.Println("result Huruf")
	fmt.Println(result)
	
	fmt.Println("result angka")
	var value1 = 201
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}