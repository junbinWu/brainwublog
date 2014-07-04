package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Email   string  //未导出的字段，首字母是小写的
}

func main() {
	t := template.New("t")
	t, _ = t.Parse("hello {{.UserName}}! {{.Email}}")
	p := Person{UserName: "Astaxie", Email:"@gmail.com"}
	t.Execute(os.Stdout, p)
}
