package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (employee FullTimeEmployee) String() string {
	return fmt.Sprintf("\nid: %d, name: %s, age: %d ", employee.id, employee.name, employee.age)
}

func newFullTimeEmployee(name string, age int, id int) *FullTimeEmployee {
	newEmployee := FullTimeEmployee{}
	newEmployee.name = name
	newEmployee.age = age
	newEmployee.id = id

	return &newEmployee
}

func GetMessage(p Person, e Employee) {
	fmt.Printf("El Id es %d %s with age %d\n", e.id, p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Maria"
	ftEmployee.age = 27
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee.Person, ftEmployee.Employee)

	ftEmployee1 := newFullTimeEmployee("Javier Mantilla", 20, 13544171)
	fmt.Println(*ftEmployee1)
}
