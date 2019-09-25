package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

func unmarshal(filename string) (post Post, err error) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}
