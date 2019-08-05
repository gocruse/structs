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

// Introduce ...
func (s Student) Introduce() string {
	return "Hello, My name is " + s.Name + " and my class is " + s.Class
}

func (s Student) sing() string {
	return "tada di dam dam dam di da dam dam"
}
