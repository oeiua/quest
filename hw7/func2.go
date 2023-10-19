package hw7

import (
	"fmt"
	"math/rand"
	"sync"
)

var ch2 = make(chan int)
var result = make(chan int, 2)
var quit1 = make(chan int)
var min, max int

func Channels2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go range_generator(&wg)
	go range_comparer()
	wg.Wait()
}

func range_generator(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		val := rand.Intn(20)
		fmt.Printf("value %v\n", val)
		ch2 <- val
	}
	quit1 <- 0
	fmt.Printf("min: %v\n", <-result)
	fmt.Printf("max: %v\n", <-result)
}

func range_comparer() {
	for i := 0; ; i++ {
		select {
		case val := <-ch2:
			fmt.Printf("compare %v\n", i)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}

		case <-quit1:
			fmt.Printf("exit\n")
			result <- min
			result <- max
			return
		}
	}
}
