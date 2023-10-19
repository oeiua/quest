package hw6_1

import "fmt"

type Train struct {
	speed int
}

func (train *Train) Move() {
	fmt.Printf(">> Train is moving\n")
}
func (train *Train) Stop() {
	fmt.Printf("Train is stopped\n")
}
func (train *Train) ChangeSpeed(speed int) {
	train.speed = speed
	fmt.Printf("Train speed is %v\n", speed)
}

func (train *Train) OutPassenger() {
	fmt.Printf("<- Train take out passenger\n")
}
func (train *Train) InPassenger() {
	fmt.Printf("-> Train take in passenger\n")
}