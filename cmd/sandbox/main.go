package main

import (
	"fmt"
)

// formatting example
type mystruct struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	x := 2
	y := 3
	z := x<<8 + y<<16
	fmt.Printf("%#x\n", z) // 0x30200
	// In C++ 0x10000000
	// https://wandbox.org/permlink/1FjZrX2rtT7bbWTb

	t := mystruct{name: "hello", value: 42}
	fmt.Print(t)

	obj := Object{owner: "chris"}
	user := "bob"

	owner := obj.Owner()
	if owner != user {
		obj.SetOwner(user)
	}

	fmt.Print(obj.Owner())
}

type Object struct {
	owner string
}

func (s *Object) Owner() string {
	return s.owner
}

func (s *Object) SetOwner(o string) {
	s.owner = o
}
