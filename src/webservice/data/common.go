package data

import (
	"database/sql"
	// PostgreSQLのドライバをインポートする。
	_ "github.com/lib/pq"
)

// 
var Db *sql.DB

func init() {
	var err error
	// DBをオープンする。
	// 第1引数は、データベースドライバの名前を指定する。
	// 第2引数は、データソース名を指定する。
	Db, err = sql.Open("postgres", "host=web_api_postgres user=gwp password=gwp dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}
