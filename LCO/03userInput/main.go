package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating  for our Pizza: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

}
