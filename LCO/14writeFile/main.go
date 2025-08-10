package main

import (
	"fmt"
	"os"
)

func main() {
	//create a new file (or overwrites if exist)
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//writes data
	bytesWritten, err := file.WriteString("Hello file I/O in Golang")
	if err != nil {
		fmt.Println("Error writting to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file \n", bytesWritten)
}
