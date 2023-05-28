package main

import "fmt"

func main() {
	iniArray := [...]int{1,2,3,4,5}
	iniSllice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSllice)


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

	newSlice := make([]string, 3, 5)

	newSlice[0] = "Lulung"
	newSlice[1] = "Satrio"
	newSlice[2] = "Prayuda"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)



}