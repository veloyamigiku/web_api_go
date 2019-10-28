package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	body, err := searchRepositoryOrderStars("ios")
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
	repositories, err := parseSearchRepository(body)
	if err != nil {
		panic(err)
	}
	for _, repository := range repositories {
		// strconv.Itoaは、数値を文字列に変換する。
		fmt.Println(strconv.Itoa(repository.ID) + ":" + repository.Name)
	}

}
