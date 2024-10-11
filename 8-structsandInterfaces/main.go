package main

import "fmt"

// import "fmt"

// func main() {
// 	var s Speaker

// 	// p := Person{Name: "Alice"}
// 	// s = p
// 	s = Person{Name: "Alice"}

// 	s.Speak()
// }

// type Speaker interface {
// 	Speak()
// }
// type Person struct {
// 	Name string
// }

// func (p Person) Speak() {
// 	fmt.Println("Hello, I am ", p.Name)
// }

type Speaker interface {
	Speak()
}
type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println("Heloo I am ", p.Name)
}

type Dog struct {
	Breed string
}

func (d Dog) Speak() {
	fmt.Println("Woof! I am a ", d.Breed)
}

func main() {
	p := Person{Name: "Alice"}

	d := Dog{Breed: "Labrodor"}

	var s Speaker
	s = p
	s.Speak()

	s = d
	s.Speak()

}
