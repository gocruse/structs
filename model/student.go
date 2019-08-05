package model

import "fmt"

// Student ...
type Student struct {
	Name   CustomString
	Class  string
	Quotes DeretanKata
}

// CustomString  must set to public (first letter must capitalized) so it can used accross projects
type CustomString string
type Siswa Student

type DeretanKata []string

func (d DeretanKata) Cetak() {
	for _, item := range d {
		fmt.Printf("\"%s\"\n", item)
	}
}
