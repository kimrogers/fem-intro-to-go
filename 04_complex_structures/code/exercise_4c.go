package main

import "fmt"

// Part 1
func Average(scores []float64) float64 {
	total := 0.0
	for _, value := range scores {
		total += value
	}
	return total / float64(len(scores))
}

// Part 2
func existsInMap(pet string) bool {

	petMap := map[string]string{
		"muffin": "cat",
		"disco":  "dog",
		"sunny":  "dog",
	}
	// for key := range petMap {
	// 	// fmt.Println(key, value)
	// 	if key == pet {
	// 		return true
	// 	}
	// }
	_, ok := petMap[pet]
	// fmt.Println("pet one", value, ok)
	return ok
}

// Part 3
func addGroceries(newStuff ...string) []string {
	var groceries = []string{"bread", "milk", "cheese", "detergent"}
	for _, value := range newStuff {
		groceries = append(groceries, value)
	}
	return groceries

}

func main() {

	var scoresArray = [5]float64{4, 5, 6.5, 7, 3}
	// fmt.Println(scoresArray)
	fmt.Println(Average(scoresArray[:]))

	fmt.Println(existsInMap("john"))
	fmt.Println(existsInMap("disco"))

	fmt.Println(addGroceries("soap", "toothpaste"))

}
