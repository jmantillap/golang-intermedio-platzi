package main

import "fmt"

type Empleado struct {
	id   int
	name string
}

func main() {

	e := Empleado{}
	fmt.Printf("%v", e)

	e.id = 1
	e.name = "Name"
	fmt.Printf("%v", e)
}
