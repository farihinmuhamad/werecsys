// "id": "we_2_2_b_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Energy Monitor",
// "title": "Check Energy Peak Level",
// "description": "Check if generator energy output is Optimal Peak, Standard Level, or Low Energy using integer thresholds."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'HighEnergy' of type int and assign it a value
	const HighEnergy int = 18000
	// TODO: Declare an integer variable 'currentEnergy' and assign it a value
	var currentEnergy int = 14000
	// TODO: Create an if statement to check if 'currentEnergy' is greater than or equal to 'HighEnergy'
	if currentEnergy >= HighEnergy {
		// TODO: Print the "Optimal Peak" level message
		fmt.Println("Optimal Peak")
		// TODO: Else if statement to check if 'currentEnergy' is greater than 9000
	} else if currentEnergy > 9000 {
		// TODO: Print the "Standard Level" message
		fmt.Println("Standard Level")
		// TODO: Else statement to handle low energy situations
	} else {
		// TODO: Print the "Low Energy" level message if none of the above conditions are met
		fmt.Println("Low Energy")
	}
}
