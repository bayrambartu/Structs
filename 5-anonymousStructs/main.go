package main

import "fmt"

func main() {
	person := struct {
		Name string
		Age  int
	}{"Eve", 29}

	fmt.Println(person)

}
