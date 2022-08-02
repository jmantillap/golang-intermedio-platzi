package main

import "fmt"

type Empleado struct {
	id   int
	name string
}

func (e *Empleado) SetId(id int) {
	e.id = id
}

func (e *Empleado) SetName(name string) {
	e.name = name
}

func (e *Empleado) GetId() int {
	return e.id
}

func (e *Empleado) GetName() string {
	return e.name
}

func main() {

	e := Empleado{}
	fmt.Printf("%v\n", e)

	e.id = 1
	e.name = "Name"
	fmt.Printf("%v\n", e)

	e.SetId(10)
	e.SetName("Javier")
	fmt.Printf("%v\n", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
