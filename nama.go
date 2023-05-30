package main

type Nama struct {
	namaPertamaT  string
	namaTerakhirT string
}

func GetNama(a, b string) Nama {
	return Nama{
		namaPertamaT:  namaPertama(a),
		namaTerakhirT: namaAkhir(b),
	}
}

func namaPertama(a string) string {
	return a
}

func namaAkhir(b string) string {
	return b
}