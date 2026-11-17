// "id": "we_1_3_e_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Thermodynamics",
// "title": "Boiler Temperature Safety Check",
// "description": "Check if the boiler temperature readings are within safe limits. Declare temperature readings, define a max limit using a constant, and evaluate safety status."
package main

import "fmt"

func main() {
	// TODO: Declare a constant int 'maxTempLimit' and assign it a value
	const maxTempLimit int = 153
	// TODO: Declare a string variable 'coolantState' and assign it a value
	var coolantState string = "On"
	// TODO: Declare an int variable 'measuredTemp' using Type Inference
	measuredTemp := 133
	// TODO: Print safety status checking if measuredTemp is less than maxTempLimit and coolantState is "On"
	fmt.Println(measuredTemp < maxTempLimit && coolantState == "On")
}