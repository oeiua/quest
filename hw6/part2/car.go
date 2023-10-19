package hw6_1

import "fmt"

type Car struct {
	speed int
}

func (car *Car) Move() {
	fmt.Printf(">> Car is moving\n")
}
func (car *Car) Stop() {
	car.speed = 0
	fmt.Printf("Car is stopped\n")
}
func (car *Car) ChangeSpeed(speed int) {
	car.speed = speed
	fmt.Printf("Car speed is %v\n", speed)
}

func (car *Car) OutPassenger() {
	fmt.Printf("<- Car take out passenger\n")
}
func (car *Car) InPassenger() {
	fmt.Printf("-> Car take in passenger\n")
}
