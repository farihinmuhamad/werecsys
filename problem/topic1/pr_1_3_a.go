// "id": "pr_1_3_a",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Flight Gate Check",
// "title": "Boarding Pass Verification",
// "description": "Verify passenger boarding details using multi variable declaration of strings, comparison operators, and printing the status.",
// "answer_options": "var passengerClass, assignedGate string = "Premium", "Gate-A9"; expectedGate := "Gate-A9"; isCorrectGate := assignedGate == expectedGate; fmt.Println(passengerClass, isCorrectGate); var passengerClass string = "Premium"; iscorrectGate := expectedGate +assignedGate"
package main

import "fmt"

func main() {
	// TODO: Declare Multi Variables 'passengerClass' and 'assignedGate' and assign it a value
	var passengerClass, assignedGate string = "Premium", "Gate-A9"
	// TODO: Declare a Type Inference variable 'expectedGate' (string) and assign it a value of "Gate-A9"
	expectedGate := "Gate-A9"
	// TODO: Declare a Type Inference variable 'isCorrectGate' (bool) using a comparison operator to check if 'assignedGate' is equal to 'expectedGate'
	isCorrectGate := assignedGate == expectedGate
	// TODO: Print the value of the 'passengerClass' and 'isCorrectGate' variables
	fmt.Println(passengerClass, isCorrectGate)
}