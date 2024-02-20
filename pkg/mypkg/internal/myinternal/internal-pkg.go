package myinternal

import (
	"fmt"
)

// this function can call only inside mypkg package
func SayInternal() {
	fmt.Println("Hello, From Internal Package!")
}