package practice

import "fmt"

func Sum() {
	arr := [4]int{5, 10, 15, 20}
	var total int = arr[0] + arr[1] + arr[2] + arr[3]
	fmt.Println("The sum of", arr[0:4], "is", (total))
}
