package main

import (
	"fmt"
)

func main() {
	var user struct{Name string; Gender int}
	user.Name = "dotcoo"
	user.Gender = 1
	fmt.Printf("%#v\n", user)

	var i float64 = 9.0/5
	fmt.Println(i)
}

