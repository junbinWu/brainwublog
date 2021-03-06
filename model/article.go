package model

import (
	"database/sql"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Date    string
	Author  string
	Title   string
	Content template.HTML
	Uid     string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:wjb@/BrainWu_Blog?charset=utf8")
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	if err != nil {
		panic("BrainWu_Blog DataBases Connection Failed")
	}
}

func GetArticleByUid(u int) *Article {
	row := db.QueryRow("select * from article where id = ?", u)
	var uid string
	var author string
	var date string
	var title string
	var content string
	err := row.Scan(&uid, &author, &date, &title, &content)
	if err != nil {
		return nil
	}
	article := &Article{Author:author, Date:date, Title:title, Content:template.HTML(content), Uid:uid}
	return article
}

func GetArticles(start int, limit int) []*Article {
	var articles []*Article
	//是否安全
	rows, err := db.Query("select * from article limit ?,?", start, limit)
	checkErr(err)
	for rows.Next() {
		var uid string
		var author string
		var date string
		var title string
		var content string

		err = rows.Scan(&uid, &author, &date, &title, &content)
		if err != nil {
			return nil
		}
		article := &Article{Author:author, Date:date, Title:title, Content:template.HTML(content), Uid:uid}
		articles = append(articles, article)
	}
	return articles
}

func GetArticleCounts() int {
	row := db.QueryRow("select count(*) from article")
	var count int
	row.Scan(&count)
	return count
}

func GetAllArticles() []*Article {
	var articles []*Article
	//是否安全
	rows, err := db.Query("select * from article")
	checkErr(err)
	for rows.Next() {
		var uid int
		var author string
		var date string
		var title string
		var content string

		err = rows.Scan(&uid, &author, &date, &title, &content)
		if err != nil {
			return nil
		}
		article := &Article{Author:author, Date:date, Title:title, Content:template.HTML(content)}
		articles = append(articles, article)
	}
	return articles
}

func SaveArticle(article Article) {
	stmt, err := db.Prepare("insert article set author=?,title=?,content=?")
	checkErr(err)
	_, err = stmt.Exec(article.Author, article.Title, string(article.Content))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
