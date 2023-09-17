package main

import (
	"fmt"
	"internal/pkg/quest/internal/pkg/quest"
)

func main() {

	fmt.Printf("Start you journey from here\n")
	fmt.Printf("f/b/r/l - to move forward, back, right, left\n")
	fmt.Printf("p - to pickup item from ground\n")
	fmt.Printf("o - to look around\n")
	fmt.Printf("s - to get status of hero\n")

	var action string
	hero := quest.InitHero()
	maze := quest.InitMaze(&hero)
	game := true
	for game {

		fmt.Scanf("%s", &action)

		switch action {
		case "left", "l":
			if maze.CanMove("left") {
				maze = quest.Move(maze)
				break
			}
		case "right", "r":
			if maze.CanMove("right") {
				maze = quest.Move(maze)
				break
			}
		case "back", "b":
			if maze.CanMove("back") {
				maze = quest.MoveBack(maze)
				break
			}
		case "forward", "f":
			if maze.CanMove("forward") {
				maze = quest.Move(maze)
				break
			}
		case "pickup", "p":
			if maze.Item != nil {
				quest.PickupItem(maze)
			}
			break
		case "status", "s":
			hero.GetStats()
			break
		case "observe", "o":
			quest.Observe(&maze)
			break
		case "suicide":
			game = false
		}

		action = ""

		if !hero.IsAlive() {
			game = false
		}
	}
	fmt.Print("Game over!")
}
