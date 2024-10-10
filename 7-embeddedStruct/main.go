package main

import "fmt"

type Animal struct {
	Name string
}
type Dog struct {
	Animal
	Breed string
}

func (a Animal) Speak() {
	fmt.Println("Animal is making a sound")
}
func main() {
	d := Dog{Animal: Animal{Name: "Fido"}, Breed: "Labrador"}
	fmt.Println(d.Name)
	d.Speak()

}
