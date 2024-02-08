package main

import "fmt"

// Ref https://www.youtube.com/watch?v=m1Uy0WQ2Xns&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=4

func main() {
	// How to print
	fmt.Println("Hello World")

	// How to declare a variable
	var name string = "John" // var name = "John"
	fmt.Println(name)
	num := 15
	fmt.Println(num)
	// Data type
	// int int8 int16 int32 int64 => this is int but in many range int 8 is 8 bit and so on
	// bool => true or false
	// string => "Hello World"
	// uint uint8 uint16 uint32 uint64 => this is int but only positive number
	// float32 float64 => this is float but in many range float 32 is 32 bit and so on. This is like float and double in C
	// rune => equal to int32 and represent a Unicode code point
	// complex64 complex128 => this is complex number.

	// How to declare a constant
	const pi = 3.14
	fmt.Println(pi)
	const (
		monday = 1
		tuesday = 2
	)
	fmt.Println(monday, tuesday)

	// in go we can use printf to format the output
	fmt.Printf("Hello %s\n", name)
	// %T is to print the type of the variable
	// %v is to print the value of the variable
	fmt.Printf("Type: %T Value: %v\n", num, num);

	// Type Conversion
	var a int = 10
	var b float64 = float64(a)
	fmt.Printf("Type: %T Value: %v\n", b, b)

	// Type Inference
	var infer1 int
	infer2 := infer1
	fmt.Printf("Type: %T Value: %v\n", infer2, infer2)
	infer2 = 10
	fmt.Printf("Type: %T Value: %v\n", infer2, infer2)
	// infer2 is changed but infer1 is not changed
	fmt.Printf("Type: %T Value: %v\n", infer1, infer1)

	// Shift Operator
	// We can use >> and << to shift the bit
	// 1 << 10 is 1024
	shift := 1 << 10
	fmt.Printf("Type: %T Value: %v\n", shift, shift)
}
