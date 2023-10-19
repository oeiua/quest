package hw6

import (
	"fmt"
)

func sortingCenter(posts []Sortable) {

	for _, item := range posts {
		if _, i := item.(*Letter); i {
			fmt.Print("This letter goes to mailman. ")
			item.Sort()
		}
		if _, i := item.(*Package); i {
			fmt.Print("This package goes to courier. ")
			item.Sort()
		}
	}
}

func PostOffice() {
	postOffice := []Sortable{
		&Letter{From: "Anton", To: "Greg"},
		&Package{From: "Rodion", To: "Alex"},
		&Package{From: "Iliya", To: "Greg"},
		&Letter{From: "Greg", To: "Rodion"},
	}
	sortingCenter(postOffice)
}
