package practice

import (
	"fmt"
)

func HouseSearch1(name string) {

	Household := [8]string{"Jim", "Daisy", "Kelsey", "Jenna", "Rob", "Ted", "Maria", "Nate"}
	for _, value := range Household {
		if value == name {
			fmt.Printf("this person lives here")
			return
		}
	}
	if name == "Daisy" {

		fmt.Println("They don`t live here, but may come in")
	} else {
		fmt.Println("Stranger Danger!")
	}

}

// func main() {

// 	HouseSearch("Rob")
// }
