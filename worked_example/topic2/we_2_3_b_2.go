// "id": "we_2_3_b_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Database Health Monitor",
// "title": "Database Connection Alert",
// "description": "Determine alert levels based on active connections and environment type using nested conditional logic."
package main

import "fmt"

func main() {
	// TODO: Declare three integer variables 'activeConnections', 'highAlertLimit', and 'regAlertLimit' using multi variable declaration and assign them values
	var activeConnections, highAlertLimit, regAlertLimit int = 15, 25000, 50000
	// TODO: Declare a boolean variable 'isProduction' and assign it a value
	var isProduction bool = true
	// TODO: Create an if statement to check if 'isProduction' is true
	if isProduction {
		// TODO: Create an if statement to check if 'activeConnections' is greater than 10
		if activeConnections > 10 {
			// TODO: Calculate the alert severity for high active connections in production and print the result
			alertSeverity := highAlertLimit * activeConnections
			fmt.Println("Production alert level:", alertSeverity)
		} else {
			// TODO: Calculate the alert severity for regular active connections in production and print the result
			alertSeverity := regAlertLimit * activeConnections
			fmt.Println("Production regular level:", alertSeverity)
		}
		// TODO: Create an else statement to handle the case when 'isProduction' is false
	} else {
		// TODO: Create an if statement to check if 'activeConnections' is greater than 15
		if activeConnections > 15 {
			// TODO: Calculate the alert severity for high active connections in staging and print the result
			alertSeverity := highAlertLimit * activeConnections
			fmt.Println("Staging alert level:", alertSeverity)
		} else {
			// TODO: Calculate the alert severity for regular active connections in staging and print the result
			alertSeverity := regAlertLimit * activeConnections
			fmt.Println("Staging regular level:", alertSeverity)
		}
	}
}
