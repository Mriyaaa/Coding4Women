package practice

import "fmt"

func MyMap(arr []string) map[string]int {

	myMap := make(map[string]int)
	for _, v := range arr {
		_, exists := myMap[v]
		if !exists {
			myMap[v] = 1
		} else {
			myMap[v] += 1
		}
	}

	fmt.Println(myMap)
	return myMap
}

// func main() {
// 	MyFruit := []string{"apple", "banana", "strawberry", "banana"}
// 	MyMap(MyFruit)
// }