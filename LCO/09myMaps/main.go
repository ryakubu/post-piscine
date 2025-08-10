package main

import "fmt"

func main() {
	fmt.Println("Maps n golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops are intersting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)

		fmt.Printf("For Key V, Value is %v\n", value)
	}
}
