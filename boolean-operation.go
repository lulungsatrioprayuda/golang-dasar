package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	var lulusNilaiAkhir1 bool = nilaiAkhir > 80
	var lulusAbsensi1 bool = nilaiAbsensi > 80
	
	var lulusNilaiAkhir2 bool = nilaiAkhir > 80
	var lulusAbsensi2 bool = nilaiAbsensi >= 80

	var lulus1 bool = lulusNilaiAkhir1 && lulusAbsensi1
	var lulus2 bool = lulusNilaiAkhir2 && lulusAbsensi2

	//hasil true karena dua duanya true
	fmt.Println(lulus2)

	//hasil false karena salah satu ada yang false
	fmt.Println(lulus1)
}