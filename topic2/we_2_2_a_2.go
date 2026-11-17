// "id": "we_2_2_a_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Logging System",
// "title": "Check Empty Log Message",
// "description": "Verify if a log message string is empty using an if statement. Declare a string variable, and check if it equals an empty string."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'logMsg' and assign it a value
	var logMsg string = "Start up checking complete"
	// TODO: Create an if statement to check if 'logMsg' is empty
	if logMsg == "" {
		// TODO: Print the "Empty Log" status message
		fmt.Println("Empty Log")
		// TODO: Use an else statement to handle non-empty logs
	} else {
		// TODO: Print the 'logMsg' if it is not empty if none of the above conditions are met
		fmt.Println(logMsg)
	}
}
