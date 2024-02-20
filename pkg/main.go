package main

import (
	"fmt"
	"test/mypkg"
)

// path of package can copy from go.mod file and add /package-folder-name

func main() {
	fmt.Println("Hello")
	mypkg.SayHello()
}