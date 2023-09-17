package quest

import (
	"fmt"
	"math/rand"
)

type Item struct {
	item  string
	value int
}

var itemTypes = []string{"food", "weapon"}

func InitItem() Item {
	random := rand.Intn(len(itemTypes))
	item := Item{
		item:  itemTypes[random],
		value: rand.Intn(9) + 1,
	}
	return item
}

func ObserveItem(item *Item) {
	if item.item == "food" {
		fmt.Printf("You see %v(%v) on the floor\n", item.item, item.value)
	}
	if item.item == "weapon" {
		fmt.Printf("You see %v with %v damage on the floor\n", item.item, item.value)
	}
}
