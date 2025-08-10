// Goroutines
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, msg)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	go printMessage("Goroutine")
	printMessage("Main")

}
