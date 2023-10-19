package hw6

import "fmt"

type Letter struct {
	From string
	To   string
}

func (letter *Letter) Sort() {
	fmt.Printf("This letter from %v to %v\n", letter.From, letter.To)
}
