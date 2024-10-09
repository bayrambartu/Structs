// package main

// import "fmt"

// func main() {
// 	var customer1 = Customer{id: 12, name: "bartu"}

// 	customer1.age = 13
// 	fmt.Println(customer1)

// }

// type Customer struct {
// 	id   int64
// 	name string
// 	age  int
// }

// package main

// import "fmt"

// func main() {
// 	var customer1 = Customer{id: 12, name: "bartu", age: 21, adress: Adress{city: "Istanbul", district: "Uskudar"}}
// 	var customer2 = Customer{id: 2, name: "bayram", age: 34}

// 	customer1.age = 31
// 	fmt.Println(customer1)
// 	fmt.Println(customer2)

// }

// type Customer struct {
// 	id     int64
// 	name   string
// 	age    int
// 	adress Adress
// }

// type Adress struct {
// 	city     string
// 	district string
// }

// package main

// import "fmt"

// func main() {
// 	var customer1 = Customer{id: 12, name: "bartu", age: 21}
// 	customer1.ToString()

// }

// type Customer struct {
// 	id   int64
// 	name string
// 	age  int
// 	// adress Adress
// }

// // type Adress struct {
// // 	city     string
// // 	district string
// // }

// func (customer Customer) ToString() {

// 	fmt.Printf("Name: %s, Age: %d", customer.name, customer.age)

// }

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 21}

	fmt.Println(p)
}
