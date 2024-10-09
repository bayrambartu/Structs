package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// func (p *Person) HaveBirthday() {
// 	p.Age += 1
// 	fmt.Println(p)
// }

func (p Person) HaveBirthday() {
	p.Age += 1
	fmt.Println(p)

}

func main() {
	p := Person{Name: "Bob", Age: 25}
	p.HaveBirthday()
	fmt.Println(p)
}
