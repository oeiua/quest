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

func Move(maze Maze) Maze {
	newMaze := RandMaze(&maze)
	newMaze.Hero = maze.Hero
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

func PickupItem(maze Maze) {
	if maze.Item.item == "food" {
		maze.Hero.health = maze.Hero.health + maze.Item.value
	}
	if maze.Item.item == "weapon" {
		maze.Hero.damage = maze.Item.value
	}
	maze.Item = nil
}
