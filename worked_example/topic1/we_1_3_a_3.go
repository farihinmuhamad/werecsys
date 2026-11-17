// "id": "we_1_3_a_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Inventory Storage",
// "title": "Storage Section Verification",
// "description": "Verify storage section assignment using multi variable declaration of strings, comparison operators, and printing the status."
package main

import "fmt"

func main() {
	// TODO: Declare Multi Variables 'itemCategory' and 'warehouseSection' (int) and assign it a value
	var itemCategory, warehouseSection int = 1, 3
	// TODO: Declare a Type Inference variable 'assignedSection' (int) and assign it a value
	assignedSection := 3
	// TODO: Declare a Type Inference variable 'isCorrectSection' (bool) using a comparison operator to check if 'warehouseSection' is equal to 'assignedSection'
	isCorrectSection := warehouseSection == assignedSection
	// TODO: Print the value of the 'itemCategory' and 'isCorrectSection' variables
	fmt.Println(itemCategory, isCorrectSection)
}