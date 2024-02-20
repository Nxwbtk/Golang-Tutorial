package main

import "fmt"

type Student struct {
	fname string
	lname string
}

type Teacher struct {
	nickName string
}

type Action interface {
	Sleep() string
}

func (std Student) Sleep() string {
	return (std.fname + " is sleeping")
}

func (teacher Teacher) Sleep() string {
	return (teacher.nickName + " is sleeping")
}

func makeAction (a Action) {
	fmt.Println(a.Sleep())
}

func main() {
	std := Student {
		fname: "New",
		lname: "Test",
	}
	t := Teacher {
		nickName: "Teacher A",
	}
	// struct if it implement all function in interface it will count as that interface
	makeAction(std)
	makeAction(t)
}