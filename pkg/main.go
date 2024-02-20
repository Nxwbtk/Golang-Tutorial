package main

import (
	"fmt"
	"test/mypkg"
)

// path of package can copy from go.mod file and add /package-folder-name
// to call pkg must use name of package and follow with function name ps. folder name can be anything but must import as same as folder name.
// 1 level can have only 1 package meaning everyfile in same folder must have same package name

func main() {
	fmt.Println("Hello")
	mypkg.SayHello()
}