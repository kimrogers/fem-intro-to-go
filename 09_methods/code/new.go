package main

import "fmt"

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	statement := fmt.Sprintf("%s \r\n You have been assigned to table %d. Your table is %s, exactly %f.2 meters from here. \n You will be sitting next to %s.", "Kim", table, direction, distance, neighbor)
	return statement
}

func main() {
	fmt.Printf(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
