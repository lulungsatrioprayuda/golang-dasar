package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[0:11]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))



	var slice2 = bulan[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Alung")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)
}