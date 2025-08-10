package main

import (
	"fmt"
	"os"
)

func main() {
	//Read from file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("File Contents")
	fmt.Println(string(data))
}
