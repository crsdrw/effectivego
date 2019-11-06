package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// formatting example
type mystruct struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	//	operatorTest()
	//	setterGetterTest()
	err := readFileTest()
	//err := readFileByteByByteTest()
	if err != nil {
		log.Printf("Failed readFileTest: %v", err)
	}

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

func operatorTest() {
	x := 2
	y := 3
	z := x<<8 + y<<16
	log.Printf("%#x", z) // 0x30200
	// In C++ 0x10000000
	// https://wandbox.org/permlink/1FjZrX2rtT7bbWTb
}

func setterGetterTest() {
	t := mystruct{name: "hello", value: 42}
	fmt.Println(t)

	obj := Object{owner: "chris"}
	user := "bob"

	owner := obj.Owner()
	if owner != user {
		obj.SetOwner(user)
	}

	fmt.Println(obj.Owner())
}

func readFileTest() error {
	f, err := os.Open("./testdata/test.txt")
	if err != nil {
		return err
	}
	buf := make([]byte, 100)
	n, err := f.Read(buf[70:100])
	log.Printf("read: %v bytes", n)
	log.Printf("bytes: %v", buf[0:100])
	if err == io.EOF {
		log.Printf("found EOF")
		return nil
	}
	return err
}

func readFileByteByByteTest() error {
	f, err := os.Open("./testdata/test.txt")
	if err != nil {
		return err
	}
	buf := make([]byte, 100)
	n, err := readFileByteByByte(buf, f)
	log.Printf("read byte-by-byte: %v bytes", n)
	log.Printf("bytes: %v", buf[0:n])
	if err == io.EOF {
		log.Printf("found EOF")
		return nil
	}
	return err
}

func readFileByteByByte(buf []byte, f io.Reader) (int, error) {
	var n int
	var err error
	for i := 0; i < 32; i++ {
		nbytes, e := f.Read(buf[i : i+1]) // Read one byte.
		n += nbytes
		if nbytes == 0 || e != nil {
			err = e
			break
		}
	}
	return n, err
}
