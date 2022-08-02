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
	endDate string
}

// Composicion sobre herencia
func (ft FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a full time employee", ft.name, ft.age)
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (te TemporaryEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a temprary employee", te.name, te.age)
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func GetMessage(pi PrintInfo) {
	fmt.Println(pi.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
	GetMessage(ftEmployee)

}
