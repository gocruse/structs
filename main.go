package main

import (
	"fmt"

	"github.com/gocruse/structs/model"
)

func main() {
	// this should be throw error
	stdntDoe := model.Student{"John Doe", "D3B", 10}
	fmt.Println(stdntDoe)

	stdntSnow := model.Student{
		Name:  "John Snow",
		Class: "D3A",
		age:   10, // this will throw error
	}
	fmt.Println(stdntSnow)

	// this should be error too, because accessing the private struct
	einstein := model.teacher{
		Name: "Einstein",
		Age:  65,
	}

	fmt.Println(einstein)
}
