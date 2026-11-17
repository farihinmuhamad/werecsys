// "id": "we_2_2_b_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Signal Strength",
// "title": "Evaluate Connection Quality",
// "description": "Assess wireless network signal strength and print connection quality using float64 constant and variable thresholds."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'MinSignal' of type float64 and assign it a value
	const MinSignal float64 = 479.5
	// TODO: Declare a float64 variable 'currentSignal' and assign it a value
	var currentSignal float64 = 321.4
	// TODO: Create an if statement to check if 'currentSignal' is greater than or equal to 'MinSignal'
	if currentSignal >= MinSignal {
		// TODO: Print the "Excellent Connection" status message
		fmt.Println("Excellent Connection")
		// TODO: Else if statement to check if 'currentSignal' is greater than 417.8
	} else if currentSignal > 417.8 {
		// TODO: Print the "Fair Connection" status message
		fmt.Println("Fair Connection")
		// TODO: Else statement to handle poor connections
	} else {
		// TODO: Print the "Poor Connection" status message if none of the above conditions are met
		fmt.Println("Poor Connection")
	}
}
