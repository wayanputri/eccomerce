package data

import (
	"belajar/bareng/features"
	"strconv"
)

func ConversiHarga(harga []string) string {
	var hargatotal int
	for _, hargatot := range harga {
		hargaInt, err := strconv.Atoi(hargatot)
		if err != nil {
			return ""
		}
		hargatotal += hargaInt
	}

	hargaString := strconv.Itoa(hargatotal)
	return hargaString
}

func AppendHarga(transaction []features.Transaction)[]string{
	var harga []string
	for _, frice := range transaction {
		harga = append(harga, frice.TotalHarga)
	}
	return harga
}
