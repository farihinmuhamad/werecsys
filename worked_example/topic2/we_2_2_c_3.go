// "id": "we_2_2_c_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "User Access Verification",
// "title": "Check Account Access Priority",
// "description": "Verify user entry attempts and failure rates. Note: deliberate differences applied — uses separate variable declarations and inverts logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'MaxAttemptsLimit' of type int and assign it a value
	const MaxAttemptsLimit int = 5
	// TODO: Declare integer variable 'entryAttempts' and assign it a value
	var entryAttempts int = 3
	// TODO: Declare integer variable 'failureRate' and assign it a value
	var failureRate int = 15
	// TODO: Create an if statement to check if 'entryAttempts' is greater than 'MaxAttemptsLimit' AND 'failureRate' is greater than 20
	if entryAttempts > MaxAttemptsLimit && failureRate > 20 {
		// TODO: Print the "Block Account" action message
		fmt.Println("Block Account")
		// TODO: Create an else-if statement to check if 'entryAttempts' is greater than 2 OR 'failureRate' is greater than 10
	} else if entryAttempts > 2 || failureRate > 10 {
		// TODO: Print the "Verify Account" action message
		fmt.Println("Verify Account")
		// TODO: Else statement to handle normal access
	} else {
		// TODO: Print the "Allow Account" status message if none of the above conditions are met
		fmt.Println("Allow Account")
	}
}
