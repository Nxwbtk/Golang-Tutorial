package mypkg

import (
	"fmt"
)

// public function must start with capital letter
func SayHello() {
	fmt.Println("Hello, From Package!")
}

// private function must start with small letter
func sayPV() {
	fmt.Println("Hello, From PV!")
}