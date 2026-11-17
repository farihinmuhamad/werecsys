// "id": "we_1_2_e_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Retail Stock",
// "title": "Warehouse Reorder Alert",
// "description": "Check warehouse retail stock levels. Declare stock item counts, define min thresholds using multi variable declaration, and evaluate status using logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'stockItemA' and assign it a value
	var stockItemA int = 36
	// TODO: Declare an integer variable 'stockItemB' and assign it a value
	var stockItemB int = 68
	// TODO: Declare two integer variables 'minThresholdA' and 'minThresholdB' using multi variable declaration and assign them values
	minThresholdA, minThresholdB := 86, 66
	// TODO: Declare a Type Inference variable 'needsReorder' (bool) using logical AND operator to check if stock items are less than thresholds
	needsReorder := (stockItemA < minThresholdA) && (stockItemB < minThresholdB)
	// TODO: Print the value of needsReorder variable
	fmt.Println("Needs reorder:", needsReorder)
}