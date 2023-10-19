package hw6_1

func Travel() {
	var routeStart = Route{name: "Lviv"}
	var car = Car{speed: 0}
	routeStart.Add(&car)
	car.InPassenger()
	car.Move()
	car.ChangeSpeed(60)
	
	var routeSecond = Route{name: "Kyiv"}
	routeSecond.Add(&car)
	car.Stop()
	car.OutPassenger()

	var train = Train{speed: 0}
	routeSecond.Add(&train)
	train.InPassenger()
	train.Move()
	train.ChangeSpeed(120)

	var routeThird = Route{name: "Dnipro"}
	routeThird.Add(&train)
	train.Stop()
	train.OutPassenger()

	var plane = Plane{speed: 0}
	routeThird.Add(&plane)
	plane.InPassenger()
	plane.Move()
	plane.ChangeSpeed(420)

	var routeFinish = Route{name: "Odessa"}
	plane.Stop()
	routeFinish.Add(&plane)
	plane.OutPassenger()

}
