package main

import "fmt"

// In go array type is fixed size
// declare as var arr[5] int

// Slice is like array but it can change the size or we can say it's dynamic array
// declare as var slice[] int

// Map is like dictionary in python or json in javascript
// declare as var m map[string]int
// string is key and int is value

// Struct is the same as struct in C
// declare as:
// type person struct {
// 	name string
// 	age int
// }

func printArr(arr [5]int) {
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println("Array at element ", i, " is ", arr[i])
	}
}

func printSlice(slice []int) {
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println("Slice at element ", i, " is ", slice[i])
	}
}

func printMap (m map[string]int) {
	fmt.Println(m)
	// This is equivalent to python for i in m: but m has range
	for key, value := range m {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}

func main() {
	var arr[5] int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	printArr(arr)

	// Slice
	slice := []int {1, 2, 3, 4, 5}
	printSlice(slice)

	// Map
	m := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}
	printMap(m)

	type Person struct {
		name string
		age int
	}
	p := Person{name: "John", age: 25}
	fmt.Println(p)
}