package main

import (
	"fmt"
	"github.com/gocruse/model"
)



func main() {
	stdntDoe := model.Student{"John Doe", "D3B"}
	fmt.Println(stdntDoe)

	stdntSnow := model.Student{
		Name:  "John Snow",
		Class: "D3A",
	}
	fmt.Println(stdntSnow)
}
