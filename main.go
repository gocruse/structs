package main

import (
	"fmt"

	"github.com/gocruse/structs/model"
)

func main() {
	deretanKata := model.DeretanKata{"Ora Et Labora", "Que Sera Sera"}
	deretanKata.Cetak()
	stdntSnow := model.Student{
		Name:   model.CustomString("John Snow"),
		Class:  "D3A",
		Quotes: deretanKata,
	}
	fmt.Println(stdntSnow)
	siswa := model.Siswa(stdntSnow)
	fmt.Println(siswa)
}
