// "id": "we_2_3_d_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Emergency Response Protocol",
// "title": "Evaluate Emergency Actions",
// "description": "Assess incident warning levels and print safety actions using a switch statement with fallthrough. Note: deliberate differences applied — switch without expression and explicit variable declaration."
package main

import "fmt"

func main() {
	// TODO: Declare string variable 'hazardSeverity' and assign it a value
	var hazardSeverity string = "Red"
	// TODO: Use switch statement without expression to evaluate hazardSeverity and print corresponding safety actions with fallthrough.
	switch {
	// TODO: Case statement to check if hazardSeverity is equal to "Red"
	case hazardSeverity == "Red":
		fmt.Println("Evacuate Sector")
		// TODO: Use fallthrough to continue to the next case
		fallthrough
	// TODO: Case statement to check if hazardSeverity is equal to "Orange"
	case hazardSeverity == "Orange":
		fmt.Println("Shutdown Reactor")
		fallthrough
	// TODO: Case statement to check if hazardSeverity is equal to "Yellow"
	case hazardSeverity == "Yellow":
		fmt.Println("Deploy Safety Shield")
	// TODO: Use default for determining Default Case for other severity levels
	default:
		fmt.Println("Status Clear")
	}
}