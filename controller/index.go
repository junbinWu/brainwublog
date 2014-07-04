package controller

import (
	"web"
	"model"
	"util"
	"html/template"
	"fmt"
	"strconv"
)

const (
	PageItemCount int = 5
)

//Data too long for column 'content' at row 1
func Index(ctx *web.Context) {
	t := util.ParseFiles("view/header.tmpl", "view/index.html", "view/nav.tmpl")
	type PageIndex struct {
		Link  template.URL
		Index int
	}
	var Data struct{
			Articles []*model.Article
			Indexs   []PageIndex
		}
	start, err1 := strconv.Atoi(ctx.Params["start"])
	limit, err2 := strconv.Atoi(ctx.Params["limit"])
	if err1 == nil && err2 == nil {
		Data.Articles = model.GetArticles(start, limit)
	} else {
		Data.Articles = model.GetArticles(0, PageItemCount)
	}
	for _, v := range Data.Articles {
		contentRune := []rune(string(v.Content))
		if len(contentRune) > 600 {
			contentRune = contentRune[:600]
			v.Content = template.HTML(string(contentRune)+"......")
		}
	}
	//做分页
	count := model.GetArticleCounts()
	var pages int
	var mode int = -1
	if count%PageItemCount == 0 {
		pages = count/PageItemCount
	} else {
		pages = count/PageItemCount+1
		mode = count%PageItemCount
	}
	for i := 0; i < pages ; i++ {
		s := i * PageItemCount
		var l int
		if i == pages-1 && mode != -1 {
			l = mode
		} else {
			l = PageItemCount
		}
		Data.Indexs = append(Data.Indexs, PageIndex{Index:i+1, Link:template.URL("/index?start="+strconv.Itoa(s)+"&limit="+strconv.Itoa(l))})
	}
	fmt.Println(t.ExecuteTemplate(ctx.ResponseWriter, "body", Data))
}

func Favicon(ctx *web.Context) {
}

//mysql 生成uuid
func Article(ctx *web.Context) {
	var art *model.Article
	p := ctx.Request.URL.Path
	p = p[9:]
	if i, err := strconv.Atoi(p); err != nil {
		ctx.Redirect(302,"/index")
	} else {
		art = model.GetArticleByUid(i)
	}
	ctx.WriteString(string(art.Content))
}
