// "id": "we_1_3_a_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Car Dealership",
// "title": "Vehicle Model Check",
// "description": "Check vehicle inventory model using string declarations, comparison operators, and printing the status."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'vehicleType' and assign it a value
	var vehicleType string = "Sedan"
	// TODO: Declare a string variable 'modelName' and assign it a value
	var modelName string = "Model-S"
	// TODO: Declare a Type Inference variable 'targetModel' (string) and assign it a value
	targetModel := "Model-S"
	// TODO: Print 'vehicleType' and the result of checking if 'modelName' is equal to 'targetModel'
	fmt.Println(vehicleType, modelName == targetModel)
}