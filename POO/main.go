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

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

func (fT FullTimeEmployee) getMessage() string {
	return "Empleado tiempo completo"
}
func (fTemp TemporatyEmployee) getMessage() string {
	return "Empleado temporal"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Printf(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Nombre"
	ftEmployee.id = 40

	fmt.Printf("%v \n", ftEmployee)
	// tempEmployee := TemporatyEmployee{}

	getMessage(ftEmployee)

}
