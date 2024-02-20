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

type LoginError struct {
	uname string
	msg string
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login Error")
}

func login(uname, pass string) error {
	if (uname == "" || pass == "") {
		return &LoginError{uname: uname, msg: "Invalid"}
	}
	return errors.New("No Still no")
}

// There is no try catch in go so we will use error (error is an interface)
func main () {
	// fmt.Println(divide(10, 2))
	// fmt.Println(divide(10, 0))
	err := login("", "")
	if (err != nil) {
		switch e := err.(type) {
		case *LoginError:
			fmt.Println("Error", e.msg)
		default:
			fmt.Println(e)
		}
	}
}