package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=web_api_postgres user=gwp password=gwp dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post {}
	err = Db.QueryRow("select id, content, author from posts where id =$1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}
