// "id": "we_1_2_b_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Scaling",
// "title": "Integer Scaling Calculation",
// "description": "Calculate a scaled result using a constant factor and integer scale value."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'BaseFactor' of type int and assign it a value
	const BaseFactor int = 227
	// TODO: Declare a Type Inference variable 'scaleMultiplier' (int) and assign it a value
	scaleMultiplier := 121
	// TODO: Calculate Type Inference variable 'scaledResult' using formula: 2 * BaseFactor * scaleMultiplier
	scaledResult := 2 * BaseFactor * scaleMultiplier
	// TODO: Print the value of the 'scaledResult' variable
	fmt.Println(scaledResult)
}