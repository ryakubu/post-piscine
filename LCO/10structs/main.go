package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 18}
	fmt.Println(hitesh)

	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)

}
