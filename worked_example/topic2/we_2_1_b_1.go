// "id": "we_2_1_b_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "System Authentication",
// "title": "Pin Validation Check",
// "description": "Verify if the entered security pin matches the registered pin. Declare expected and entered pin values as integers, then check equality to determine access status."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'expectedPin' and assign it a value
	expectedPin := 9841
	// TODO: Declare a variable 'enteredPin' of type int and assign it a value
	var enteredPin int = 9840
	// TODO: Create an if statement to check if 'enteredPin' is equal to 'expectedPin'
	if enteredPin == expectedPin {
		// TODO: Inside the if block, print the value of "Pin Correct"
		fmt.Println("Pin Correct")
		// TODO: Add an else block
	} else {
		// TODO: Inside the else block, print the value of "Pin Incorrect"
		fmt.Println("Pin Incorrect")
	}
}
