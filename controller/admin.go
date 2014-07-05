package controller

import (
	"web"
	"html/template"
	"net/http"
	"time"
	"model"
)

func Login(ctx *web.Context) {
	name := ctx.Params["username"]
	pass := ctx.Params["password"]

	if isValid(name, pass) {
		coo := &http.Cookie{Name:"admin", Value:"WuJunBin", Expires:time.Now().AddDate(0, 0, 1)}
		ctx.SetCookie(coo)
		t, _ := template.ParseFiles("view/editor.html")
		t.Execute(ctx.ResponseWriter, nil)
	} else {
		t, _ := template.ParseFiles("view/login.html")
		t.Execute(ctx.ResponseWriter, nil)
	}
}

func isValid(arg0 string, arg1 string) bool {
	if len(arg0) == 0 || arg0 != "WuJunBin" {
		return false
	}
	if len(arg1) == 0 || arg1 != "BrainWu" {
		return false
	}
	return true
}

func Admin(ctx *web.Context) {
	if cookie, err := ctx.GetCookie("admin") ; err != nil {
		ctx.Redirect(302, "/login")
	} else {
		if cookie.Value == "WuJunBin" {
			t, _ := template.ParseFiles("view/editor.html")
			t.Execute(ctx.ResponseWriter, nil)
		} else {
			ctx.Redirect(302, "/login")
		}
	}
}

//加 标题 时间 等参数
func Editor(ctx *web.Context) {
	var editTv string = ctx.Params["editor1"]
	var editTitle string = ctx.Params["title"]
	article := model.Article{
		Author : "BrainWu",
		Title  :editTitle,
		Content : template.HTML(editTv),
	}
	model.SaveArticle(article)
	ctx.Redirect(302,"/index")
}
