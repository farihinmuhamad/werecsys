// "id": "we_2_3_a_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Pressure Sensor Evaluation",
// "title": "Check Pressure Sensor Status",
// "description": "Assess pressure levels using a float64 variable and output the pressure state using if-else statements."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference float64 variable 'measuredPressure' and assign it a value
	measuredPressure := 225.5
	// TODO: Create an if statement to check if 'measuredPressure' is greater than 285.0
	if measuredPressure > 285.0 {
		// TODO: Print the Critical Pressure status message
		fmt.Println("Critical Pressure")
		// TODO: Create an else-if statement to check if 'measuredPressure' is greater than 195.0
	} else if measuredPressure > 195.0 {
		// TODO: Print the High Pressure status message
		fmt.Println("High Pressure")
		// TODO: Create another else-if statement to check if 'measuredPressure' is greater than 115.0
	} else if measuredPressure > 115.0 {
		// TODO: Print the Moderate Pressure status message
		fmt.Println("Moderate Pressure")
		// TODO: Create an else statement to print Low Pressure
	} else {
		// TODO: Print the Low Pressure status message if none of the conditions are met
		fmt.Println("Low Pressure")
	}
}
