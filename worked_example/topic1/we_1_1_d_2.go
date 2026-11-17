// "id": "we_1_1_d_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Access Rights",
// "title": "Negating Availability Status",
// "description": "Example showing how to determine restriction and blocked status using logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare boolean variable 'isAvailable' and assign it true
	isAvailable := true
	// TODO: Declare boolean variable 'isRestricted' and assign it false
	isRestricted := false
	// TODO: Determine 'isBlocked' status by negating 'isAvailable' or checking 'isRestricted'
	isBlocked := !isAvailable || isRestricted
	// TODO: Print the value of the 'isBlocked' variable
	fmt.Println(isBlocked)
}