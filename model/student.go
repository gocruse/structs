package model

import "fmt"

// Student ...
type Student struct {
	Name          CustomString
	Class         string
	Quotes        DeretanKata
	SelfIntroduce Introduction
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

// IntroduceSelf is a Student function
func (s Student) IntroduceSelf() string {
	return s.SelfIntroduce(s)
}

// Introduction is a custom type of a function.
type Introduction func(s Student) string
