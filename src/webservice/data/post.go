package data

// 送受信するJSONに対する構造体（投稿）。
type Post struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func Retrieve(id int) (post Post, err error) {
	post = Post {}
	// DB.QueryRowは、SQLクエリを実行する。
	// 第1引数は、SQLクエリの文字列を指定する。
	// 第2引数以降は、SQLクエリの変数($1,$2,..)に割り当てるデータを指定する。
	// Row.Scanは、結果を引数のポインタに割り当てる。
	err = Db.QueryRow("select id, content, author from posts where id =$1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// 構造体Postのメソッド。
// テーブルpostsに投稿を登録する。
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// DB.Prepareは、引数のSQLクエリを元にStmt（プリペアードステートメント）を作成する。
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

func (post *Post) Update() (err error) {
	// DB.Execは、指定のSQLクエリを実行する。（行結果は返却しない）
	// 第1引数は、SQLクエリを指定する。
	// 第2引数以降は、SQLクエリ内の変数に割り当てる値を指定する。
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.ID)
	return
}
