package mypkg

import (
	"fmt"
)

// public function must start with capital letter
func SayHello() {
	fmt.Println("Hello, From Package!")
	// call private function
	sayPV()
}

// private function must start with small letter
// but you can call private function from public function
func sayPV() {
	fmt.Println("Hello, From PV!")
}