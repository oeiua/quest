package hw6

import "fmt"

type Package struct {
	From string
	To   string
}

func (pack *Package) Sort() {
	fmt.Printf("This package from %v to %v\n", pack.From, pack.To)
}
