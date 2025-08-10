package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"-"`
	Completed bool   `json:"completed"`
}

func main() {
	/*url := "http://jsonplaceholder.typicode.com/todos/1/"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {*/
	todoItem := &Todo{1, 1, "Release video on Go", false}

	/*decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&todoItem); err != nil {
		log.Fatal("Decode error: ", err)
	}*/

	// convert back to json
	todo, err := json.MarshalIndent(todoItem, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(todo))
}
