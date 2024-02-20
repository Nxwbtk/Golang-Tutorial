package main

import "fmt"

type Linkedlist struct {
	data int
	next *Linkedlist
}

type Method interface {
	setNext(*Linkedlist)
}

func (obj *Linkedlist) setNext (newObj *Linkedlist) {
	obj.next = newObj
}

func makelist(oldObj *Linkedlist, newObj *Linkedlist) {
	oldObj.setNext(newObj)
}

func main () {
	fmt.Println("Hello Liked list")
	obj1 := &Linkedlist{
		data: 10,
		next: nil,
	}

	obj2 := &Linkedlist{
		data: 20,
		next: nil,
	}
	makelist(obj1, obj2)
	var tmp *Linkedlist
	tmp = obj1
	for tmp != nil {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}