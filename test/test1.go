package main

import (
	"fmt"
	"bytes"
)

type Person struct {
	Name string
	Age int
	Gender int
}

func (this *Person) say() {
	fmt.Println(this.Name, " Hello!")
}

type Student struct {
	Person
	School string
}


func main() {
	s := Student{Person{Name:"BrainWu",Age:21,Gender:0}, "beijingshifandaxue"}
	s.say()
	testU(s)
	fmt.Println(s)

	s1 := []int{1,2,3,4,5}
	testS(s1)
	fmt.Println(s1)

	str := []string{"wujunbin"}
	testStr(str)
	fmt.Println(str)

	fmt.Println(string(bytes.Map(func(r rune) rune {
			if r == '-' {
				r = ':'
			}
			return r
		}, []byte("wu-ju-nb-in"))))
	var a = bytes.Replace([]byte("wu-ju-nb-in"),[]byte("-"),[]byte(":"),-1)
	fmt.Println(string(a))
}

func testU(u Student) {
	u.Name = "--------------"
	fmt.Println(u)
}

func testS(s1 []int) {
	s1[0] = 100
	fmt.Println(s1)
}

func testStr(s []string) {
//	rs := []rune(s)
//	rs[0] = 'l'
//	fmt.Println(s)
	s[0] = "--==="
	fmt.Println(s)
}
