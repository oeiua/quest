package hw7

import (
	"fmt"
	"math/rand"
)

var ch = make(chan int)
var out = make(chan int)
var quit = make(chan int)
var avg = 0

func Channels1() {
	go generator()
	go average()
	go output()
	<-quit
}

func generator() {
	for i := 0; i < 20; i++ {
		var rnd = rand.Int()
		ch <- rnd
	}
	quit <- avg
}

func average() {
	for i := 1; ; i++ {
		val := <-ch
		avg = (avg + val) / i
		out <- avg
	}
}

func output() {
	for i := 0; ; i++ {
		fmt.Println(<-out)
	}
}
