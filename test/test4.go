package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go testG(i)
	}
	time.Sleep(10 * time.Second)
}

var t time.Time

func testG(s int) {
	fmt.Println(s)
	fmt.Println("before ", s)
	time.Sleep(time.Second)
	fmt.Println(s)
}

