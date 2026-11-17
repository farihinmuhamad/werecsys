// "id": "we_2_1_a_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Financial Approval",
// "title": "Credit Limit Eligibility Check",
// "description": "Determine if a client qualifies for a credit limit adjustment based on their current score. Declare a constant limit threshold and a variable score, then check if it meets the criteria."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'reqCredit' of type float64 and assign it a value
	const reqCredit float64 = 680.5
	// TODO: Declare a variable 'clientCredit' of type float64 and assign it a value
	var clientCredit float64 = 715.0
	// TODO: Create an if statement to check if 'clientCredit' is greater than 'reqCredit'
	if clientCredit > reqCredit {
		// TODO: Inside the if block, print the value of "Eligible"
		fmt.Println("Eligible")
	}
}
