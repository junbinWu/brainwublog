package model

import (
	"testing"
	"fmt"
	"html/template"
)

func Test_GetArticles(t *testing.T) {
	as := GetArticles(0, 3)
	if len(as) <= 0 {
		fmt.Println(as)
		fmt.Println(len(as))
		t.Fail()
	}
}

//这种应该如何写测试用例
func Test_SaveArtivle(t *testing.T) {
	article := Article{
		Date:"2014-06-06",
		Author :"BrainWu",
		Title  :"Test",
		Content :template.HTML("Test_SaveArticle"),
	}
	SaveArticle(article)
}

func Test_GetArticleCounts(t *testing.T) {
	if count := GetArticleCounts(); count > 0 {
		t.Log(count)
	} else {
		t.Error("Test_GetArticleCounts Error")
	}
}

func Test_GetArticleByUid(t *testing.T) {
	if art := GetArticleByUid(2); art == nil {
		t.Fail()
	}
}

//create table article (
//uid int(10) not null auto_increment,
//author varchar(64) default null,
//date varchar(64) default null,
//title varchar(64) default null,
//content text,
//primary key (uid)
//) CHARSET=utf8;
