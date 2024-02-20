package main

import "fmt"

// We can use struct in struct
type SmallStruct struct {
	smmsg string
}

type BigStruct struct {
	msg	string
	small SmallStruct
}

type Student struct {
	name string
	age	int
	grade float32
}
func main() {
	var std1 Student
	std1.name = "John"
	std1.age = 25
	std1.grade = 3.5
	fmt.Println(std1)
	fmt.Println("Name: ", std1.name)

	// We can declare as array of struct
	var std2 [2]Student
	std2[0].name = "John"
	std2[0].age = 25
	std2[0].grade = 3.5

	std2[1].name = "Doe"
	std2[1].age = 26
	std2[1].grade = 3.6
	fmt.Println(std2)

	for i := 0; i < 2; i++ {
		fmt.Printf("Student No. %d Name: %s Age: %d Grade %.2f \n", i + 1, std2[i].name, std2[i].age, std2[i].grade)
	}

	var nestedStruct BigStruct

	nestedStruct.msg = "Hello Big Struct"
	nestedStruct.small.smmsg = "Hello Small struct"
	fmt.Println("Nested struct: ", nestedStruct)
	fmt.Println("Big struct message: ", nestedStruct.msg)
	fmt.Println("Small struct message: ", nestedStruct.small.smmsg)
}