package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the file name
	const filename = "complete_example.txt"

	// 1. Create and write to the file

	initialContent := []byte("This is the first line of text.\n")
	err := os.WriteFile(filename, initialContent, 0644)
	if err != nil {
		fmt.Printf("Error creating and writing to file: %v\n", err)
		return
	}
	fmt.Println("Step 1: Successfully created and wrote initial content to the file.")

	// Ensure the file is removed when the program exits. This is good practice for cleanup.
	defer os.Remove(filename)

	// 2. Read the file back
	fmt.Println("\nStep 2: Reading the file for the first time...")
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n---\n%s---\n", content)

	// 3. Append a new line to the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file for append: %v\n", err)
		return
	}
	// Always remember to close the file to release the resource.
	defer file.Close()

	// Write the new content to the file.
	newLine := "This is a new line appended to the file.\n"
	_, err = file.WriteString(newLine)
	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}
	fmt.Println("Step 3: Successfully appended a new line to the file.")

	// 4. Read the file again to see the appended content
	fmt.Println("\nStep 4: Reading the file again to show the appended content...")
	updatedContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading updated file: %v\n", err)
		return
	}
	fmt.Printf("Updated file content:\n---\n%s---\n", updatedContent)
}
