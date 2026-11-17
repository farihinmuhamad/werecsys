// "id": "pr_2_3_d",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Industrial Process Safety",
// "title": "Hazard Alert Progression",
// "description": "Determine safety mitigation actions using a switch statement with fallthrough to print all subsequent action steps for a given alert level progression. The program should evaluate the 'severityLevel' string variable and print the corresponding safety actions based on the alert level: 'Red' for 'Evacuate Zone', 'Orange' for 'Power Down', 'Yellow' for 'Deploy Shield', and 'All Clear' for any other value. Use fallthrough to ensure that all subsequent actions are printed for higher severity levels.",
// "answer_options": "switch severityLevel {; case "Red":; fallthrough; default:; fallthrough next case:; default severityLevel:"
package main

import "fmt"

func main() {
	severityLevel := "Red"
	// TODO: Use switch statement to evaluate 'severityLevel' and print corresponding safety mitigation actions with fallthrough.
	switch severityLevel {
	// TODO: Case statement to check if 'severityLevel' is equal to "Red"
	case "Red":
		fmt.Println("Evacuate Zone")
		// TODO: Use fallthrough to continue to the next case
		fallthrough
	case "Orange":
		fmt.Println("Power Down")
		fallthrough
	case "Yellow":
		fmt.Println("Deploy Shield")
	// TODO: Use default for determining Default Case for other severity levels
	default:
		fmt.Println("All Clear")
	}
}