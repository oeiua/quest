package quest

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	left     bool
	right    bool
	forward  bool
	back     bool
	previous *Maze
	Hero     *Hero
	Item     *Item
}

func InitMaze(hero *Hero) Maze {
	var maze = Maze{
		left:     false,
		right:    false,
		forward:  true,
		back:     false,
		Hero:     hero,
		previous: nil,
		Item:     nil,
	}
	return maze
}

func RandMaze(maze *Maze) Maze {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	item := InitItem()
	return Maze{
		left:     r.Intn(2) == 1,
		right:    r.Intn(2) == 1,
		forward:  r.Intn(2) == 1,
		back:     true,
		Hero:     nil,
		previous: maze,
		Item:     &item,
	}
}

func Observe(maze *Maze) {
	if maze.left {
		fmt.Printf("You can move left. ")
	}
	if maze.right {
		fmt.Printf("You can move right. ")
	}
	if maze.forward {
		fmt.Printf("You can move forward. ")
	}
	if maze.back {
		fmt.Printf("You can move back. ")
	}
	fmt.Printf("\n")
	if maze.Item != nil {
		ObserveItem(maze.Item)
	}
}

func MoveLeft(maze Maze) Maze {
	newMaze := RandMaze(&maze)
	newMaze.Hero = maze.Hero
	fmt.Printf("You move left\n")
	return newMaze
}

func MoveRight(maze Maze) Maze {
	newMaze := RandMaze(&maze)
	newMaze.Hero = maze.Hero
	fmt.Printf("You move right\n")
	return newMaze
}

func MoveForward(maze Maze) Maze {
	newMaze := RandMaze(&maze)
	newMaze.Hero = maze.Hero
	fmt.Printf("You move forward\n")
	return newMaze
}

func MoveBack(maze Maze) Maze {
	newMaze := maze.previous
	newMaze.Hero = maze.Hero
	maze.Hero = nil
	return *newMaze
}

func (maze *Maze) CanMove(side string) bool {
	if side == "right" && maze.right {
		return true
	}
	if side == "left" && maze.left {
		return true
	}
	if side == "forward" && maze.forward {
		return true
	}
	if side == "back" && maze.back {
		return true
	}
	fmt.Printf("You cant move %v\n", side)
	return false
}

func PickupItem(maze Maze) Maze {
	if maze.Item.item == "food" {
		maze.Hero.health = maze.Hero.health + maze.Item.value
		fmt.Printf("You pick up %v(%v)\n", maze.Item.item, maze.Item.value)
	}
	if maze.Item.item == "weapon" {
		maze.Hero.damage = maze.Item.value
		fmt.Printf("You pick up %v with %v damage\n", maze.Item.item, maze.Item.value)
	}
	maze.Item = nil
	return maze
}

func StartGame(){
	fmt.Printf("Start you journey from here\n")
	fmt.Printf("f/b/r/l - to move forward, back, right, left\n")
	fmt.Printf("p - to pickup item from ground\n")
	fmt.Printf("o - to look around\n")
	fmt.Printf("s - to get status of hero\n")

	var action string
	hero := InitHero()
	maze := InitMaze(&hero)
	game := true
	for game {

		fmt.Scanf("%s", &action)

		switch action {
		case "left", "l":
			if maze.CanMove("left") {
				maze = MoveLeft(maze)
			}
		case "right", "r":
			if maze.CanMove("right") {
				maze = MoveRight(maze)
			}
		case "back", "b":
			if maze.CanMove("back") {
				maze = MoveBack(maze)
			}
		case "forward", "f":
			if maze.CanMove("forward") {
				maze = MoveForward(maze)
			}
		case "pickup", "p":
			if maze.Item != nil {
				maze = PickupItem(maze)
			}
		case "status", "s":
			hero.GetStats()
		case "observe", "o":
			Observe(&maze)
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