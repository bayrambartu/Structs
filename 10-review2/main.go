package main

import "fmt"

type car struct {
	make  string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	bedSize int
}

func main() {
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)
}
