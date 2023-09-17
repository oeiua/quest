package quest

import "fmt"

type Hero struct {
	health int
	damage int
}

func InitHero() Hero {
	var hero = Hero{
		damage: 1,
		health: 10,
	}

	return hero
}

func (hero *Hero) GetStats() {
	fmt.Printf("Health: %vhp | ", hero.health)
	fmt.Printf("Damage: %v\n", hero.damage)
}

func (hero *Hero) IsAlive() bool {
	return hero.health > 0
}
