// "id": "we_3_1_b_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Connection Test",
// "title": "Network Connection Diagnostic",
// "description": "Simulate a diagnostic connection test that logs ping alerts and breaks early on success."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'maxRetryAttempts' and assign it a value
	var maxRetryAttempts int = 5
	// TODO: Declare a string variable 'pingAlertMsg' and assign it a value
	var pingAlertMsg string = "Ping Server Attempt"
	// TODO: Create a regular for loop to log server ping attempts and break early
	for i := 1; i <= maxRetryAttempts; i++ {
		// TODO: Print the 'pingAlertMsg' to the console
		fmt.Println(pingAlertMsg)
		// TODO: Create an if statement to check if 'i' is equal to 3
		if i == 3 {
			// TODO: Print "Connection Restored" to the console
			fmt.Println("Connection Restored")
			// TODO: Break the loop if
			break
		}
	}
}