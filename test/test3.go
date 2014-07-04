package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Now()

	y, m, d := tm.Date()
	fmt.Println(y)
	fmt.Println(int(m))
	fmt.Println(d)
	fmt.Println(tm)
	fmt.Println(tm.YearDay())
	fmt.Println(tm.Location())
	fmt.Println(tm.ISOWeek())

	addT := tm.AddDate(1,1,1)
	fmt.Println(addT)
}
