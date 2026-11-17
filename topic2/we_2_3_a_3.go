// "id": "we_2_3_a_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Water Temperature Sensor classification",
// "title": "Check Water Temperature Range",
// "description": "Assess water temperature using an integer variable. Note: deliberate differences applied — uses switch instead of if-else, and checks lower range using less-than comparisons."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference integer variable 'waterTemp' and assign it a value
	waterTemp := 225
	// TODO: Use switch statement without an expression to evaluate the temperature
	switch {
	// TODO: Case statement to check if 'waterTemp' is less than 115
	case waterTemp < 115:
		// TODO: Print the Ice State status message
		fmt.Println("Ice State")
	// TODO: Case statement to check if 'waterTemp' is less than 195
	case waterTemp < 195:
		// TODO: Print the Cold Water status message
		fmt.Println("Cold Water")
	// TODO: Case statement to check if 'waterTemp' is less than 285
	case waterTemp < 285:
		// TODO: Print the Warm Water status message
		fmt.Println("Warm Water")
	default:
		// TODO: Default case to print Boiling Water status message
		fmt.Println("Boiling Water")
	}
}
