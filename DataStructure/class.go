package main

import "fmt"

// function with return type
func getValue() int {
	return 10
}

type Student struct {
	FirstName string
	LastName string
}

// To create method of the object or in other language we call class
// We have to create receiver method
func (std Student) getFirstName() string {
	return std.FirstName
}

func (std Student) getLastName() string {
	return std.LastName
}

func (std Student) setFirstName(firstName string) {
	std.FirstName = firstName
}

func (std Student) setLastName(lastName string) {
	std.LastName = lastName
}

func main() {
	fmt.Println(getValue())
	var std Student
	// cannot set because must use pointer
	std.setFirstName("John")
	std.setLastName("Doe")
	fmt.Printf("Student name = %s %s\n", std.getFirstName(), std.getLastName())
}