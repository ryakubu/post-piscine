package main

import (
	"encoding/json"
	"fmt"
	"myapp/mathutils"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{Name: "Alice", Age: 30}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling", err)
		return
	}
	fmt.Println(string(jsonData))

	sum := mathutils.Add(100, 3)
	product := mathutils.Multiply(100, 3)

	fmt.Printf("The sum is: %v\n ", sum)
	fmt.Printf("The product is: %v\n ", product)
}
