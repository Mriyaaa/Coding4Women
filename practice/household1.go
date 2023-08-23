package practice

import (
	"fmt"
)

func HouseSearch(name string) {

	Household := [8]string{"Jim", "Daisy", "Kelsey", "Jenna", "Rob", "Ted", "Maria", "Nate"}
	for _, value := range Household {
		if value == name {
			fmt.Printf("this person lives here")

		} else {
			fmt.Println("Stranger Danger!")
		}
	}
}

// func main() {

// 	HouseSearch("Jim")

// }
