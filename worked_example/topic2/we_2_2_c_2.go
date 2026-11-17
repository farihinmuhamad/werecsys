// "id": "we_2_2_c_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Water Quality Classification",
// "title": "Check Water Condition Level",
// "description": "Assess pH level and water turbidity using integer parameters and determine the safety level."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'MaxPhLimit' of type int and assign it a value
	const MaxPhLimit int = 9
	// TODO: Declare two integer variables 'phLevel' and 'turbidity' using type inference and assign them values
	phLevel, turbidity := 7, 12
	// TODO: Create an if statement to check if 'phLevel' is greater than 'MaxPhLimit' OR 'turbidity' is greater than 25
	if phLevel > MaxPhLimit || turbidity > 25 {
		// TODO: Print the "Hazardous Water" warning message
		fmt.Println("Hazardous Water")
		// TODO: Create an else-if statement to check if 'phLevel' is greater than 6 AND 'turbidity' is greater than 10
	} else if phLevel > 6 && turbidity > 10 {
		// TODO: Print the "Warning Level" warning message
		fmt.Println("Warning Level")
		// TODO: Else statement to handle safe water status
	} else {
		// TODO: Print the "Safe Water" status message if none of the above conditions are met
		fmt.Println("Safe Water")
	}
}
