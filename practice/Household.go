package practice


func livesWithMe(person string, Household [8]string) bool {
	for _, member := range Household {
		if person == member {
			return true
		}
	}
	return false
}

// func main() {

// 	Household := [8]string{"Jim", "Daisy", "Kelsey", "Jenna", "Rob", "Ted", "Maria", "Nate"}

// 	person := "Daisy"

// 	if livesWithMe(person, Household) {
// 		fmt.Println("Daisy lives here")
// 	} else if person == "Daisy" {
// 		fmt.Println("Daisy does not live here but may come in")
// 	} else {
// 		fmt.Println("Stranger Danger!")
// 	}
//}
