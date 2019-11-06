package main

import (
	"fmt"
	"time"
)

func count(c chan int) {
	for i := 0; i < 10; i++ {
		//		fmt.Println("Counting", i)
		c <- i
		fmt.Println("Counted", i)
		time.Sleep(1 * time.Second)
	}
	//close(c)
}

func print(c chan int) {
	for i := 0; i < 10; i++ {
		//time.Sleep(1 * time.Second)
		//fmt.Println("Printing", i)
		v := <-c
		fmt.Println("Printed", v)
	}
}

func main() {
	c := make(chan int)
	go count(c)
	go print(c)
	time.Sleep(15 * time.Second)
}
