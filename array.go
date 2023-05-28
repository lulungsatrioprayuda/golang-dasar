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

	//menghitung jumlah lengt atau panjang indexnya
	fmt.Println(len(names))
	fmt.Println(len(values))

	//coba hitung array yang sudah di setting index nya namun tidak ada isinya 
	var lagi [10]string
	//hasilnya 10
	fmt.Println(len(lagi))
}