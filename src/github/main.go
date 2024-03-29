package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	// リポジトリ検索を実行する。
	body, err := searchRepositoryOrderStars("ios")
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
	
	// リポジトリ検索結果（JSON文字列）をパースする。
	repositories, err := parseSearchRepository(body)
	if err != nil {
		panic(err)
	}
	for _, repository := range repositories {
		// strconv.Itoaは、数値を文字列に変換する。
		fmt.Println(strconv.Itoa(repository.ID) + ":" + repository.Name)
	}

	// リポジトリ検索結果（JSONオブジェクト）をJSON文字列に変換する。
	fmt.Println(marshalSearchRepository(repositories))

}
