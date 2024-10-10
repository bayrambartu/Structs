package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Charlie", Age: 40}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))
}
