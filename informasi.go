package main

type Info struct {
	NamaPertama string
	NamaTengah  string
	NamaAkhir   string
	Umur        int
}

func infoAll(namaLengkap Info, umur int) Info {
	namaLengkap.Umur = umur
	return namaLengkap

}

func getNamaLengkap(a, b, c string) Info {
	return Info{
		NamaPertama: a,
		NamaTengah:  b,
		NamaAkhir:   c,
	}
}

func getUmurs(d, e int) int {
	return d * e
}