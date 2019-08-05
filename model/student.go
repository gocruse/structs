package model

// Student ...
type Student struct {
	Name  string // Public properties
	Class string // Public properties
	age   int    // private properties
}

// teacher is private propertes
type teacher struct {
	Name string
	Age  int
}
