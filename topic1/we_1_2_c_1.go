// "id": "we_1_2_c_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Inventory Load",
// "title": "Warehouse Capacity",
// "description": "Calculate total capacity of warehouse packs using Type Inference, discard an unused note variable using the blank identifier, and print the results with high volume status."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'itemCount' (int) and assign it a value
	itemCount := 139
	packSize := 147
	// TODO: Calculate Type Inference variable 'totalCapacity' using formula: itemCount * packSize
	totalCapacity := itemCount * packSize
	// TODO: Declare a Type Inference variable 'batchNote' (string) and assign it a value
	batchNote := "Standard Warehouse Pack"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'batchNote' variable
	_ = batchNote
	// TODO: Declare a Type Inference variable 'isHighVolume' (bool), assign it with a logical expression (if totalCapacity is greater than 15000)
	isHighVolume := totalCapacity > 15000
	// TODO: Print the calculated totalCapacity and isHighVolume status
	fmt.Println(totalCapacity, isHighVolume)
}