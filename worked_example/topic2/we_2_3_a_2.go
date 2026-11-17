// "id": "we_2_3_a_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Smart Grid Load Check",
// "title": "Check Smart Grid Status",
// "description": "Assess power grid load using an integer variable and output the load state using if-else statements."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'gridLoad' and assign it a value
	var gridLoad int = 225
	// TODO: Create an if statement to check if 'gridLoad' is greater than or equal to 285
	if gridLoad >= 285 {
		// TODO: Print the Critical Load status message
		fmt.Println("Critical Load")
		// TODO: Create an else-if statement to check if 'gridLoad' is greater than or equal to 195
	} else if gridLoad >= 195 {
		// TODO: Print the Peak Load status message
		fmt.Println("Peak Load")
		// TODO: Create another else-if statement to check if 'gridLoad' is greater than or equal to 115
	} else if gridLoad >= 115 {
		// TODO: Print the Standard Load status message
		fmt.Println("Standard Load")
		// TODO: Create an else statement to print Light Load
	} else {
		// TODO: Print the Light Load status message if none of the conditions are met
		fmt.Println("Light Load")
	}
}
