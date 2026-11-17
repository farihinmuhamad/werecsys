// "id": "we_2_1_d_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Batch Processing",
// "title": "Check if Batch Index is Odd or Even",
// "description": "Check if a batch processing index is odd or even using modulus and a switch statement. Note: deliberate differences applied — case 1 (odd check) is placed first, and the default handles the even case."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'batchIndex' and assign it a value
	var batchIndex int = 486
	// TODO: Use switch statement to check batchIndex with modulus operation
	switch batchIndex % 2 {
	// TODO: Case statement to check for odd index (remainder 1) — difference 1: odd case placed first
	case 1:
		// TODO: Print "Odd Index" for case 1
		fmt.Println("Odd Index")
	// TODO: Default case to handle even index — difference 2: even handled as default
	default:
		// TODO: Print "Even Index" for default case
		fmt.Println("Even Index")
	}
}
