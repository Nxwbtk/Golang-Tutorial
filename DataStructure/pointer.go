package main

import "fmt"

// pointer is the same as C

type Student struct {
	fname string
	lname string
}

type Method interface {
	getFirstName() string
	getLastName() string
	setFirstName(string)
	setLastName(string)
}

func (std *Student) setFirstName(fname string) {
	std.fname = fname
}

func (std *Student) setLastName(lname string) {
	std.lname = lname
}

func (std Student) getFirstName() string {
	return std.fname
}

func (std Student) getLastName() string {
	return std.lname
}

func printFullName (m Method) {
	fmt.Printf("Firstname = %s LastName = %s\n", m.getFirstName(), m.getLastName())
}

func main () {
	fmt.Println("Hello Pointer")
	std := &Student {
		fname: "New",
		lname: "Last Name",
	}
	fmt.Println(std)
	fmt.Printf("%p\n", std)
	fmt.Println("Before set")
	printFullName(std)

	std.setFirstName("John")
	std.setLastName("Doe")
	fmt.Println("After set")
	printFullName(std)
}