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
	// we can print capasity of slice by using cap(slice)
	fmt.Println("Capasity of slice is ", cap(slice))
	// we can print slice from index 1 to 3 by using slice[1:3] || slice[1:] || slice[:3] same as python
	subSlice := slice[:3]
	fmt.Println(subSlice)
	// but subSlice cap is 5 because it's still point to slice
	fmt.Println("Capasity of subSlice is ", cap(subSlice))

	// append slice
	slice = append(slice, 10, 11, 12)
	// This is because of slice didn't allocate memory it's just pick the memory only size it's need
	fmt.Println(slice)
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

	// We can convert array to slice by using convertArrtoSlice := arr[:]
	fmt.Println("Convert array to slice: ", arr[:])
	fmt.Printf(("Type of arr: %T\n"), arr)
	fmt.Printf(("Type of arr[:] : %T\n"), arr[:])

	// Map
	// Simple map without allocate memory
	m := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Printf("Type of m: %T\n", m)
	printMap(m)

	// We can use make to allocate memory for map
	mapMake := make(map[string]int)
	mapMake["one"] = 1
	mapMake["two"] = 2
	mapMake["three"] = 3
	fmt.Printf("Type of mapMake: %T\n", mapMake)
	printMap(mapMake)
	// We can use delete to delete key from map
	delete(mapMake, "two")
	fmt.Println("Len of mapMake is ", len(mapMake))
	printMap(mapMake)
	// To get Value of Map
	value := mapMake["one"]
	fmt.Println("Value of one is ", value)
	// We can check if key is exist by using _, ok := m["one"]
	_, ok := mapMake["one"]
	fmt.Println("Is one exist: ", ok)
	_, ok = mapMake["two"]
	fmt.Println("Is two exist: ", ok)

	// Struct
	type Person struct {
		name string
		age int
	}
	p := Person{name: "John", age: 25}
	fmt.Println(p)
}