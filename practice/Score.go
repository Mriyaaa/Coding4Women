package practice

import "fmt"

func Score(points int) {
	if points >= 100 {
		fmt.Println("Winner!!!")

	} else if points >= 50 {
		fmt.Println("Not bad!!!")

	} else {
		fmt.Println("Loser!!!")
	}

}

// func main() {

// 	Score(105)

// }
