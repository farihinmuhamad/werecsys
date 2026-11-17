// "id": "we_2_1_b_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Server Authentication",
// "title": "Access Token Validation Check",
// "description": "Validate incoming client request tokens. Declare a predefined server token, then use a short-statement if initialization to check the request token."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'savedToken' using manifest typing
	var savedToken string = "auth-token-xyz"
	// TODO: Perform short statement if initialization to check if clientToken matches savedToken
	if clientToken := "auth-token-abc"; clientToken == savedToken {
		// TODO: Print validation successful message
		fmt.Println("Validation Successful")
		// TODO: Add an else block
	} else {
		// TODO: Print validation failed message
		fmt.Println("Validation Failed")
	}
}
