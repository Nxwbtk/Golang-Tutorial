package main

import "fmt"

func sampleIf (a int) {
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}
}

func sampleSwitch (a int) {
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is neither 1 nor 2")
	}
}

func samplePreProcess(a int, b int) {
	if sum := a + b; sum > 10 {
		fmt.Println("sum is greater than 10")
	} else {
		fmt.Println("sum is less than 10")
	}
}

func sampleFor () {
	for i := 0; i < 5; i++ {
		fmt.Println("For loop : ", i)
	}
}

func sampleDoWhile() {
	// there is no do while in go so we can use for loop to do that
	i := 0
	// for without condition is equal to infinite loop
	for {
		fmt.Println("Do while loop : ", i)
		i++
		if i == 5 {
			break
		}
	}
}

func sampleWhileLoop() {
	// there is no while loop in go so we can use for loop to do that
	i := 0
	// we can declare with only ending condition
	for i < 5 {
		fmt.Println("While loop : ", i)
		i++
	}
}

func main() {
	sampleIf(20)
	sampleSwitch(2)
	samplePreProcess(5, 6)
	sampleFor()
	sampleDoWhile()
	sampleWhileLoop()
}