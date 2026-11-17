// "id": "we_2_1_c_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Engine Cooling",
// "title": "Check Engine Coolant Temperature Level",
// "description": "Check if the engine coolant temperature is Overheated, Normal, or Cool based on predefined float64 thresholds. If the coolant temp is less than or equal to the cool threshold, print 'Cool'. If greater than cool but less than or equal to normal threshold, print 'Normal'. If greater than the normal threshold, print 'Overheat'."
package main

import "fmt"

func main() {
	// TODO: Declare Type Inference float64 variable 'coolLimit' and assign it a value
	coolLimit := 27.5
	// TODO: Declare Type Inference float64 variable 'normalLimit' and assign it a value
	normalLimit := 38.5
	// TODO: Declare Type Inference float64 variable 'coolantTemp' and assign it a value
	coolantTemp := 35.0
	// TODO: Create an if statement to check if coolantTemp is below 'coolLimit'
	if coolantTemp < coolLimit {
		// TODO: Print the "Cool" status message
		fmt.Println("Cool")
		// TODO: Create an else if statement to check if coolantTemp is below 'normalLimit'
	} else if coolantTemp < normalLimit {
		// TODO: Print the "Normal" status message
		fmt.Println("Normal")
		// TODO: Create an else statement to handle other other condition  
	} else {
		// TODO: Print the "Overheat" if none of the conditions are met
		fmt.Println("Overheat")
	}
}
