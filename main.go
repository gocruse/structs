package main

import (
	"fmt"

	"github.com/gocruse/structs/model"
)

func main() {
	deretanKata := model.DeretanKata{"Ora Et Labora", "Que Sera Sera"}
	deretanKata.Cetak()

	// Create model.Introduction
	selfIntro := model.Introduction(func(s model.Student) string {
		return "Hello, My name is " + string(s.Name) + " and my class is " + s.Class
	})

	stdntSnow := model.Student{
		Name:          model.CustomString("John Snow"),
		Class:         "D3A",
		Quotes:        deretanKata,
		SelfIntroduce: selfIntro,
	}
	fmt.Println(stdntSnow.IntroduceSelf())
	siswa := model.Siswa(stdntSnow)
	fmt.Println(siswa)
}
