// "id": "pr_2_2_a",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Input Validation",
// "title": "Check Empty Input",
// "description": "In this problem, you need to declare a string variable and check if the variable is empty using an if statement. If the variable is empty, print 'Empty Input'. Otherwise, print 'Valid Input'.",
// "answer_options": "var input string = ""; if input == "" {; } else {; fmt.Println("Valid Input"); const input string = "Empty"; { else }"
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'input' and assign it an empty string
	var input string = ""
	// TODO: Create an if statement to check if 'input' is empty
	if input == "" {
		fmt.Println("Empty Input")
	// TODO: Use an else statement to print "Valid Input" if 'input' is not empty
	} else {
		// TODO: Print "Valid Input" to the console
		fmt.Println("Valid Input")
	}
}