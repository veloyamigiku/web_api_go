package main

type Post struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author Author `json:author"`
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
