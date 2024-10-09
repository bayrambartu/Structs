package main

import "fmt"

type Adress struct {
	City  string
	State string
}
type Person struct {
	Name   string
	Age    int
	Adress Adress
}

func main() {
	person1 := Person{Name: "John", Age: 35, Adress: Adress{City: "NewYork", State: "NY"}}
	fmt.Println(person1)
}
