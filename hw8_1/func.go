package hw81

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var timeout int

func Run() {
	flag.IntVar(&timeout, "timeout", 3, "How long to run orders generator")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go makeOrders(ctx)
	go ordersCalculator(ctx, &wg)

	wg.Wait()
	outputOrders()
}

var orders = make(chan Order)

var ordersSum = make(map[string]int)

func generator() Order {
	var names = []string{"Anton", "Rodion", "Greg", "Illya"}
	var items = []string{"mouse", "keyboard", "hammer", "microscope", "monitor", "ups", "modem"}
	time.Sleep(300 * time.Millisecond)
	return Order{
		name:  names[rand.Intn(len((names)))],
		item:  items[rand.Intn(len(items))],
		count: rand.Intn(5),
		cost:  rand.Intn(99) + 1,
	}
}

func makeOrders(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			orders <- generator()
		}
	}
}

func ordersCalculator(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case value := <-orders:
			ordersSum[value.name] += (value.cost * value.count)
		}
	}
}

func outputOrders() {
	fmt.Println(ordersSum)
}
