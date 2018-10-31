package main

import (
	"fmt"
)

// formatting example
type t struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	x := 2
	y := 3
	z := x<<8 + y<<16
	fmt.Printf("%#x", z) // 0x30200
	// In C++ 0x10000000
	// https://wandbox.org/permlink/1FjZrX2rtT7bbWTb

}
