// "id": "we_2_2_d_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Shipping Method Selector",
// "title": "Check Shipping Service Type",
// "description": "Select the correct delivery type based on a string shipping method value using a switch statement."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'method' and assign it a value
	var method string = "Express"
	// TODO: Implement a switch statement to handle different 'method' values and print the corresponding delivery type.
	switch method {
	// TODO: Case statement to check if 'method' is equal to "Standard"
	case "Standard":
		// TODO: Print the Standard Delivery message
		fmt.Println("Standard Delivery")
	// TODO: Case statement to check if 'method' is equal to "Express"
	case "Express":
		// TODO: Print the Express Delivery message
		fmt.Println("Express Delivery")
	// TODO: Case statement to check if 'method' is equal to "Priority"
	case "Priority":
		// TODO: Print the Priority Delivery message
		fmt.Println("Priority Delivery")
	// TODO: Use default for determining Default Case for other methods
	default:
		// TODO: Print the Default Delivery message
		fmt.Println("Default Delivery")
	}
}
