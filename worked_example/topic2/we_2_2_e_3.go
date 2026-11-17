// "id": "we_2_2_e_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Server Health Monitor",
// "title": "Check Server Latency Status",
// "description": "Assess server response times and error rates. Note: deliberate differences applied — declares variables separately, and uses logical AND operators instead of OR."
package main

import "fmt"

func main() {
	// TODO: Declare integer variable 'responseTime' and assign it a value
	var responseTime int = 82
	// TODO: Declare integer variable 'errorRate' and assign it a value
	var errorRate int = 94
	// TODO: Use switch statement without an expression to evaluate the server conditions
	switch {
	// TODO: Case statement to check if 'responseTime' is greater than 93 AND 'errorRate' is greater than 93
	case responseTime > 93 && errorRate > 93:
		// TODO: Print the Critical Outage level message
		fmt.Println("Critical Outage")
	// TODO: Case statement to check if 'responseTime' is greater than 77 AND 'errorRate' is greater than 77
	case responseTime > 77 && errorRate > 77:
		// TODO: Print the Warning Condition level message
		fmt.Println("Warning Condition")
	// TODO: Case statement to check if 'responseTime' is greater than 54 AND 'errorRate' is greater than 54
	case responseTime > 54 && errorRate > 54:
		// TODO: Print the Moderate Latency level message
		fmt.Println("Moderate Latency")
	default:
		// TODO: Default case to print Optimal Status message
		fmt.Println("Optimal Status")
	}
}
