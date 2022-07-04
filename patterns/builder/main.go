package main

import "fmt"

func main() {
	compBuilder := NewComputerBuilder()
	computer := compBuilder.CPU("Core i3").RAM(8).MB("Gigabyte")
	fmt.Println(computer)

	officeCompBuilder := NewOfficeComputerBuilder()
	officeCompBuilder.RAM(4)
	officeComputer := officeCompBuilder.Build()
	fmt.Println(officeComputer)

	anotherCompBuilder := NewComputerBuilder()
	anotherCompBuilder = anotherCompBuilder.CPU("core i9").RAM(32).MB("Asus")
	director := NewDirector(anotherCompBuilder)

	compA := director.BuildComputer()
	fmt.Println(compA)

	anotherOfficeBuilder := NewOfficeComputerBuilder()
	director.SetBuilder(anotherOfficeBuilder)

	compB := director.BuildComputer()
	fmt.Println(compB)
}
