package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	lg := log.New(os.Stdout, "----------", log.Ldate|log.Ltime)
	lg.Println("--------")

	str := "wujunbin"
	println(str[0:4])

	s := "1"
	bs := []byte(s)
	println(strconv.Atoi(string(bs)))

}

