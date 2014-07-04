package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"runtime"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:wjb@/BrainWu_Blog?charset=utf8")
	if err != nil {
		panic("BrainWu_Blog DataBases Connection Failed")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0 ; i < 10; i ++ {
		go query("wujunbinsad")
		go query("a")
		go query("asd")
		time.Sleep(1 * time.Second)
	}
	time.Sleep(60 * time.Second)
}

func query(t string) {
	db2,_ := sql.Open("mysql", "root:wjb@/BrainWu_Blog?charset=utf8")
	defer db2.Close()
	row := db2.QueryRow("select * from article where title = ?", t)
	var uid int
	var author string
	var date string
	var title string
	var content string
	row.Scan(&uid, &author, &date, &title, &content)
	fmt.Println(title)
}

