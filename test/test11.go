package main

import (
	"fmt"
	"time"
)

type Stack struct {
	Stack []int
	Index int
}

func (s *Stack) Pop() {
	s.Index--
}

func (s * Stack) Push(n int) {
	if s.Index == len(s.Stack) {
		s.Stack = append(s.Stack, n)
	} else {
		s.Stack[s.Index] = n
	}
	s.Index++
}

func main() {
	throwsPanic(func() {
		panic(time.Now())
	})
	fs := []float64{1.1, 2.56, 3.3, 4.3, 5.3}
	fmt.Println(avg(fs))

	in := make([]int, 4, 6)
	fmt.Println(len(in))
	in = append(in, 2, 4, 5, 6, 7)
	fmt.Println(&in)
	fmt.Println(in)

	fmt.Println(fbnq(20))
}

func throwsPanic(f func()) {
	defer func() {
		if x := recover(); x != nil {
			if v, ok := x.(time.Time); ok {
				fmt.Println(v.Date())
			}
		}
	}()
	f()
	fmt.Println("BrainWu")
}

func avg(fs []float64) float64 {
	var f float64
	var index int = 0
	for i, v := range fs {
		f = f+v
		index = i
	}
	f = f/float64((index + 1))
	return f
}


func fbnq(n int) int {
	x, y := 1, 1
	var sum int = x
	for y < n {
		x, y = y, x+y
		sum = sum + x
	}
	return sum
}
