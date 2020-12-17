package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// an instance of our Book struct
	book := Book{Title: "Learning Go", Author: "Gopher Go"}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
