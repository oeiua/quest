package hw6_1

type Movable interface {
	Move()
	Stop()
	ChangeSpeed(int)
}

type Rideble interface {
	InPassenger()
	OutPassenger()
}
