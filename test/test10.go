package main

import (
	"fmt"
)

func main() {
	var ar []string = []string{"1","2","3"}
	test(ar)
}

func test(args ... string) {
	fmt.Println(args)
}
