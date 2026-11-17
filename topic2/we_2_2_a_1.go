// "id": "we_2_2_a_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Calibration",
// "title": "Check Calibration Offset",
// "description": "Verify if a calibration offset is exactly the default value using an if statement. Declare a float64 variable, and check if it equals 0.05."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'calibrationOffset' and assign it 0.05
	var calibrationOffset float64 = 0.05
	// TODO: Create an if statement to check if 'calibrationOffset' is equal to 0.05
	if calibrationOffset == 0.05 {
		// TODO: Print the "Default Offset" status message
		fmt.Println("Default Offset")
		// TODO: Use an else statement to handle other offsets
	} else {
		// TODO: Print the "Custom Offset" status message if none of the above conditions are met
		fmt.Println("Custom Offset")
	}
}
