package controller

import (
	"web"
	"util"
)

func About(ctx *web.Context) {
	t := util.ParseFiles("view/header.tmpl", "view/person.html","view/nav.tmpl")
	t.ExecuteTemplate(ctx.ResponseWriter,"body",nil)
}
