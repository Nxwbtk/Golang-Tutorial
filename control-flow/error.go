package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if (b == 0) {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

type Login struct {
	uname string
	pass string
	msg string
}

func (e *Login) Error() string {
	return fmt.Sprintf("Login Error")
}

func login(uname, pass string) error {
	if (uname == "" || pass == "") {
		return &Login{uname: uname, msg: "Invalid"}
	}
	return errors.New("No Still no")
}

// There is no try catch in go so we will use error (error is an interface)
func main () {
	// fmt.Println(divide(10, 2))
	// fmt.Println(divide(10, 0))
	err := login("asdsad", "dsadsa")
	if (err != nil) {
		switch e := err.(type) {
		case *Login:
			fmt.Println("Error", e.msg)
		default:
			fmt.Println(e)
		}
	}
}