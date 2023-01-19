package main

import "fmt"

func average(arg1, arg2, arg3 float64) float64 {
	total := arg1 + arg2 + arg3
	return total / 3
}

func main() {

	fmt.Println(average(3.1, 3.2, 3.3))

}
