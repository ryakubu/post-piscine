package main

import (
	"fmt"
	"os"
)

func main() {
	// Open file in append mode
	file, err := os.OpenFile("eaxmple.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append::", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Adding a new line! \n")
	if err != nil {
		fmt.Println("Error appending to file.")
		return
	}
	fmt.Println("Successfully appended to file. ")

}
