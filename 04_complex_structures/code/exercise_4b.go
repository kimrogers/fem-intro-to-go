package main

import "fmt"

func average(nums ...float64) float64 {
	var total float64

	for _, value := range nums {
		total += value
	}

	return total / float64(len(nums))
}

func main() {

	fmt.Println(average(3.1, 3.2, 3.3))

}
