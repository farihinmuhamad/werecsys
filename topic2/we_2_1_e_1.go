// "id": "we_2_1_e_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "System Status",
// "title": "Classify System Status Code",
// "description": "Verify system health status using a switch statement on integer status codes. Declare an integer variable, and evaluate if it matches specific inactive status codes."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'statusCode' and assign it a value
	var statusCode int = 471
	// TODO: Use switch statement to check the 'statusCode' variable
	switch statusCode {
	// TODO: Case statement to check if status matches inactive codes (471 or 473)
	case 471, 473:
		// TODO: Print the "Disrupted" status message
		fmt.Println("Disrupted")
	// TODO: Use default for determining Default Case for other status codes
	default:
		// TODO: Print the "Functional" status message
		fmt.Println("Functional")
	}
}
