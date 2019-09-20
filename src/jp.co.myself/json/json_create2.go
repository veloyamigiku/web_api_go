package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments`
}

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {

	post := Post {
		ID: 1,
		Content: "Hello World",
		Author: Author {
			ID: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment {
			Comment {
				ID: 3,
				Content: "Have a great day!",
				Author: "Adam",
			},
			Comment {
				ID: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
		},
	}

	jsonFile, err := os.Create("json_create2.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
	
}