package hw6_1

import "fmt"

type Plane struct {
	speed int
}

func (plane *Plane) Move() {
	fmt.Printf(">> Plane is moving\n")
}
func (plane *Plane) Stop() {
	fmt.Printf("Plane is stopped\n")
}
func (plane *Plane) ChangeSpeed(speed int) {
	plane.speed = speed
	fmt.Printf("Plane speed is %v\n", speed)
}

func (plane *Plane) OutPassenger() {
	fmt.Printf("<- Plane take out passenger\n")
}
func (plane *Plane) InPassenger() {
	fmt.Printf("-> Plane take in passenger\n")
}