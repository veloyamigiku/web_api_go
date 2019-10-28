package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/antonholmquist/jason"
)

// Repository 構造体（リポジトリ）。
type Repository struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// リポジトリ検索結果（JSONオブジェクト）をJSON文字列に変換する。
func marshalSearchRepository(repositories []Repository) (jsonString string, err error) {
	jsonBytes, err := json.Marshal(repositories)
	if err != nil {
		return
	}
	jsonString = string(jsonBytes)
	return
}

// リポジトリ検索結果文字列（JSON）をパースして、各リポジトリのIDと名前を配列形式で返却する。
func parseSearchRepository(bodyStr string) (repositories []Repository, err error) {

	// jason.NewObjectFromByteは、JSON文字列をパースしてJSONオブジェクトを作成する。
	rootJSONObject, err := jason.NewObjectFromBytes([]byte(bodyStr))
	if err != nil {
		return
	}

	// GetObjectArrayは、項目名を指定してJSONオブジェクト内の配列を取得する。
	items, err := rootJSONObject.GetObjectArray("items")
	if err != nil {
		return
	}

	// 空のスライスを作成する。
	repositories = []Repository{}

	for _, item := range items {
		// GetInt64は、項目名を指定して整数を取得する。
		id, err := item.GetInt64("id")
		if err != nil {
			break
		}
		// GetStringは、項目名を指定して文字列を取得する。
		name, err := item.GetString("name")
		if err != nil {
			break
		}
		repository := Repository{}
		repository.ID = int(id)
		repository.Name = name
		// スライスに要素を追加して、代入する。
		repositories = append(repositories, repository)
	}

	return
}

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
