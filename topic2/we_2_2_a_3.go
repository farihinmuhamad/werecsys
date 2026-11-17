// "id": "we_2_2_a_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Session Verification",
// "title": "Check Active Session Token",
// "description": "Verify if a session token is present. Note: deliberate differences applied — uses != operator instead of ==, and swaps the if/else branch logic."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'sessionToken' and assign it an empty string
	var sessionToken string = ""
	// TODO: Create an if statement to check if 'sessionToken' is not empty (difference 1: != instead of ==)
	if sessionToken != "" {
		// TODO: Print the active session status message (difference 2: branch swap)
		fmt.Println("Active Session")
		// TODO: Use an else statement to handle empty session tokens
	} else {
		// TODO: Print the missing token status message if none of the above conditions are met
		fmt.Println("Missing Token")
	}
}
