package main

import (
	"os"
	"web"
	"controller"
	"fmt"
)

func pwd() string {
	//命令运行时所在的目录
	//比如在/etc运行的 /../../../main文件  最终获取到的 etc目录
	pwd, _ := os.Getwd()
	return pwd
}

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(pwd())
	web.AddRoute("/about", web.HandleFunc(controller.About), "GET")
	web.AddRoute("/editor", web.HandleFunc(controller.Editor), "POST")
	web.AddRoute("/admin", web.HandleFunc(controller.Admin), "GET")
	web.AddRoute("/login", web.HandleFunc(controller.Login), "GET")
	web.AddRoute("/login", web.HandleFunc(controller.Login), "POST")
	web.AddRoute("/favicon.ico", web.HandleFunc(controller.Favicon), "GET")
	web.AddRoute("/article/.*", web.HandleFunc(controller.Article), "GET")
	web.AddRoute("/(.*)", web.HandleFunc(controller.Index), "GET")
	web.Run(":8080", nil)
}
