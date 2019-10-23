package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func searchRepositoryOrderStars(q string) (bodyStr string, err error) {

	// 検索URLを作成する。
	// url.Valueは、クエリパラメータの構造体。
	v := url.Values{}
	// url.Addは、クエリパラメータを追加する。
	v.Add("q", q + " in:name")
	v.Add("sort", "starts")
	v.Add("order", "desc")
	// url.Encodeは、クエリパラメータをURLエンコードする。
	url := "https://api.github.com/search/repositories?" + v.Encode()

	// 検索URLにアクセスする。
	// http.Getは、指定のURLにアクセスして、Responseを返却する。
	res, err := http.Get(url)
	if err != nil {
		return
	}
	// Response.Bodyの型は、io.ReadCloser。
	defer res.Body.Close()
	// ioutil.ReadAllは、io.Readerを読み込んで、byteスライスで返却する。
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// string()は、byteスライスを文字列に変換する。
	bodyStr = string(body)

	return
}
