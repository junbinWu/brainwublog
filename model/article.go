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

func GetArticles(start int, limit int) []*Article {
	var articles []*Article
	//是否安全
	rows, err := db.Query("select * from article limit ?,?", start, limit)
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
	stmt, err := db.Prepare("insert article set author=?,date=?,title=?,content=?")
	checkErr(err)
	_, err = stmt.Exec(article.Author, article.Date, article.Title, string(article.Content))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
