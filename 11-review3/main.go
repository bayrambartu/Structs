// package main

// import "fmt"

// type rect struct {
// 	width  int
// 	height int
// }

// // area has a receiver of (r rect)
// // rect is the struct
// // r is the placeholder
// func (r rect) area() int {
// 	return r.width * r.height
// }
// func main() {
// 	r := rect{
// 		width:  5,
// 		height: 10,
// 	}
// 	fmt.Println(r.area())
// 	// prints 50
// }

package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization:  Basic : %s,%s", a.username, a.password)

}
func main() {
	a := authenticationInfo{
		username: "bartu",
		password: "12345",
	}

	fmt.Println(a.getBasicAuth())
}

// create the method below
