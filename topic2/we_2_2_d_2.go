// "id": "we_2_2_d_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Payment Tier Selector",
// "title": "Check Payment Tier Level",
// "description": "Select the correct membership level based on an integer payTier value using a switch statement."
package main

import "fmt"

func main() {
	// TODO: Declare type inference integer variable 'payTier' and assign it a value
	payTier := 2
	// TODO: Implement a switch statement to handle different 'payTier' values and print the corresponding access level.
	switch payTier {
	// TODO: Case statement to check if 'payTier' is equal to 1
	case 1:
		// TODO: Print the Basic Tier message
		fmt.Println("Basic Tier")
	// TODO: Case statement to check if 'payTier' is equal to 2
	case 2:
		// TODO: Print the Premium Tier message
		fmt.Println("Premium Tier")
	// TODO: Case statement to check if 'payTier' is equal to 3
	case 3:
		// TODO: Print the Enterprise Tier message
		fmt.Println("Enterprise Tier")
	// TODO: Use default for determining Default Case for other tiers
	default:
		// TODO: Print the Unknown Tier message
		fmt.Println("Unknown Tier")
	}
}
