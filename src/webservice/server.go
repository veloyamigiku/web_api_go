package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"webservice/data"
	"webservice/util"
)

func main() {

	// トークン発行のサンプルコード
	dir := util.GetCurrentDir()
	privatePath := dir + "/key/private"
	token, err := util.IssueToken(privatePath)
	if err != nil {
		panic(err)	
	}
	fmt.Println("tokenString:" + token.Token)
	publicPath := dir + "/key/public"
	res, err := util.ValidateToken(token.Token, publicPath)
	if err != nil {
		panic(err)
	}
	fmt.Print("res:")
	fmt.Println(res)
	

	// 変数の宣言＆代入。
	// Server構造体を生成する。
	server := http.Server {
		Addr: ":8080",
	}
	// http.HandleFuncは、マルチプレクサ（リクエストをハンドラ関数にリダイレクトするコード）にハンドラ関数を割り当てる。
	// http.Handleは、マルチプレクサに「ハンドラ」を割り当てる。
	// ※ハンドラは、ServeHTTPメソッドを持つ構造体を指す。
	http.HandleFunc("/post/", handleRequest)
	http.HandleFunc("/login/", handleLoginRequest)

	// HTTPサーバを作成する。
	// 第1引数は、ネットワークアドレスを指定する。
	// 第2引数は、受付ポート番号を指定する。
	server.ListenAndServe()

}

func handleLoginRequest(w http.ResponseWriter, r *http.Request) {

	// ユーザーの存在を確認する。
	r.ParseForm()
	if _, ok := r.PostForm["user"]; !ok {
		http.Error(
			w,
			"user param is not found.",
			http.StatusInternalServerError)
		return
	}
	// 外部サービス等を利用して、ユーザ名の存在を確認する。
	user := r.PostForm["user"][0]
	if user != "root" {
		http.Error(
			w,
			"user is invalie.",
			http.StatusInternalServerError)
		return
	}
	
	// 外部サービス等を利用して、ハッシュ化したユーザパスワードを取得する。


	// ユーザパスワードを検証する。


	// JWT（トークン）を発行する。
	currentDir := util.GetCurrentDir()
	privatePath := currentDir + "/key/private"
	token, err := util.IssueToken(privatePath)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	// 発行したJWT（トークン）をJSON形式で出力する。
	w.Header().Set("Content-Type", "application/json")
	jsonString, err := json.MarshalIndent(token, "", "\t")
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
	w.Write(jsonString)

}

// ハンドラ関数(引数にResponseWriterとRequestのポインタを持つ)
func handleRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		// http.Errorは、指定のエラーメッセージとステータスコードをクライアントに返却する。
		// 第1引数は、ResponseWriterを指定する。
		// 第2引数は、エラーメッセージを指定する。。
		// 第3引数は、ステータスコードを指定する。
		http.Error(
			w,
			// Errorは、エラーメッセージを返却する。
			err.Error(),
			// StatusInternalServerErrorは、ステータスコード（500）の定数。
			http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// strconv.Atoiは、引数の文字列を整数に変換する。
	// path.Baseは、引数のURL文字列をスラッシュで区切った最後の要素を返却する。
	// http.Request.URL.Pathは、リクエストURLのパス部分を返却する。
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	
	post, err := data.Retrieve(id)
	if err != nil {
		return
	}
	// json.MarshalIndentは、インターフェースをJSON形式のデータ（バイト配列）に変換する。
	// 第1引数は、変換元のインタフェースを指定する。
	// 第2引数は、(動作確認後に記載予定。)
	// 第3引数は、JSON形式のデータを整形するインデントを指定する。
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	// ResponseWriter.Headerは、Headerデータを参照する。
	// Header.Setは、ヘッダ名と値を設定する。
	w.Header().Set("Content-Type", "application/json")
	// HTTPの応答として、HTTP接続に対してバイト配列を書き込む。
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// http.Request.ContentLengthは、リクエストボディのサイズをバイト数で返却する。
	len := r.ContentLength
	// 組み込み関数makeは、スライスを作成する。
	// 第1引数は、型（[]int, []byte, etc）を指定する。
	// 第2引数は、長さを指定する。
	body := make([]byte, len)
	// http.Request.Body[io.Reader].Readは、リクエストボディを読み込んでバッファ（スライス）に保存する。
	r.Body.Read(body)
	var post data.Post
	// json.Unmarshalは、JSON形式のデータ（byte配列）を指定の型に変換する。
	json.Unmarshal(body, &post)
	err = post.Create()
	if err != nil {
		return
	}
	// ResponseWriter.WriteHeaderは、引数のステータスコードでレスポンスヘッダを送信する。
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := data.Retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := data.Retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
