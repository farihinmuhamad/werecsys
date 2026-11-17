// "id": "we_2_3_d_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Chemical Refinery Safety",
// "title": "Verify Refinery Alert Status",
// "description": "Determine refinery hazard alerts and mitigation safety levels using conditional branching."
package main

import "fmt"

func main() {
	// TODO: Declare string variable 'refineryAlert' and assign it a value
	refineryAlert := "Critical"
	// TODO: Declare string variable 'mitigationLevel' and assign it a value
	mitigationLevel := "Full Containment"
	// TODO: Create an if statement to check if refineryAlert is Critical
	if refineryAlert == "Critical" {
		// TODO: Create a nested if statement to check if mitigationLevel is Full Containment
		if mitigationLevel == "Full Containment" {
			fmt.Println("Critical Warning: Active Containment Mode")
		// TODO: Use else to determine the case when mitigationLevel is not Full Containment
		} else {
			fmt.Println("Critical Warning: Action Pending")
		}
	// TODO: Use else if to check if refineryAlert is Warning
	} else if refineryAlert == "Warning" {
		fmt.Println("Warning Status: Deploy Secondary Grid")
	// TODO: Use else to determine the case when refineryAlert is neither Critical nor Warning
	} else {
		fmt.Println("Refinery Status: Normal Green")
	}
}