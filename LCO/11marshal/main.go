// Converting from struct to Json
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 23}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error Marshaling", err)
		return
	}
	fmt.Println(string(jsonData))
}
