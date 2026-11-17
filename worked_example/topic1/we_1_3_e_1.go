// "id": "we_1_3_e_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Aquarium Monitoring",
// "title": "Water Level Safety Verification",
// "description": "Check if the aquarium water level readings are within safe limits. Declare level readings, define a safe limit using a constant, and evaluate safety status using logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 constant 'SafeLevel' and assign it a value
	const SafeLevel float64 = 75.8
	// TODO: Declare a string variable 'filterStatus' and assign it a value
	var filterStatus string = "Stopped"
	// TODO: Declare a float64 variable 'waterLevel' using Type Inference
	waterLevel := 64.2
	// TODO: Declare a Type Inference variable 'isAquariumSafe' (bool), check if waterLevel is below 'SafeLevel' AND 'filterStatus' does not equal "Stopped"
	isAquariumSafe := (SafeLevel >= waterLevel) && (filterStatus != "Stopped")
	// TODO: Print the value of 'isAquariumSafe'
	fmt.Println(isAquariumSafe)
}