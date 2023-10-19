package hw6_1

import "fmt"

type Route struct {
	Vehicles []Movable
	name     string
}

func (route *Route) Add(vehicle Movable) {
	route.Vehicles = append(route.Vehicles, vehicle)
	fmt.Printf("Arrived in %v\n", route.name)
}

func (route *Route) List() {
	for _, item := range route.Vehicles {
		fmt.Println(item)
	}
}
