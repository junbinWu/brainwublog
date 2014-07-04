package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}
//{{with .Friends}}
//{{range .}}
//my friend name is {{.Fname}}
//{{end}}
//{{end}}

func main() {
		f1 := Friend{Fname: "minux.ma"}
		f2 := Friend{Fname: "xushiwei"}
		t := template.New("fieldname example")
		t, _ = t.Parse(` hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}`)
		p := Person{UserName: "Astaxie",
			Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
			Friends: []*Friend{&f1, &f2}}
		t.Execute(os.Stdout, p)
//
//	ss := []string{"文章列表","相册","个人档案","友情链接","新浪微博","留言板"}
//	nav1 := &Nav{NavItems:ss}
//	t := template.New("ta")
//	t, _ = t.Parse(`{{range .NavItems}} {{.}} {{end}} `)
//	t.Execute(os.Stdout,nav1)
}

type Nav struct {
	NavItems []string
}
