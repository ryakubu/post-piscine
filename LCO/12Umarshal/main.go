package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	var jsonBody = `{"Title":"Golang Series", "Author":"Rinny Yaks", "Price": 123}`
	var b Book

	err := json.Unmarshal([]byte(jsonBody), &b)
	if err != nil {
		fmt.Println("Error Unmarshaling:", err)
		return
	}
	fmt.Println(b.Author)
	fmt.Println(jsonBody)
}
