// "id": "pr_1_3_b",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Warehouse Dispatch",
// "title": "Cargo Dispatch Validation",
// "description": "Evaluate if a cargo shipment is ready for dispatch using multi-variable declaration of integers, constants, blank identifier, and logical operator checks.",
// "answer_options": "const MaxWeightLimit int = 1200; var itemWeight, itemQuantity int = 83, 14; zoneCode := "Zone-B"; _ = zoneCode; hasHazardousMaterial := false; isApproved := (itemWeight * itemQuantity <= MaxWeightLimit) && !hasHazardousMaterial; fmt.Println(isApproved); var itemWeight int = 83"
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'MaxWeightLimit' and assign it a value
	const MaxWeightLimit int = 1200
	// TODO: Declare Multi Variable 'itemWeight' and 'itemQuantity' and assign it a value
	var itemWeight, itemQuantity int = 83, 14
	zoneCode := "Zone-B"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'zoneCode' variable
	_ = zoneCode
	hasHazardousMaterial := false
	// TODO: Declare a Type Inference variable 'isApproved' (bool) to check if 'itemWeight' multiplied by 'itemQuantity' is less than or equal to 'MaxWeightLimit' AND 'hasHazardousMaterial' is false
	isApproved := (itemWeight*itemQuantity <= MaxWeightLimit) && !hasHazardousMaterial
	fmt.Println(isApproved)
}