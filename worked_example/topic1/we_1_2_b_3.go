// "id": "we_1_2_b_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Investment Leverage",
// "title": "Double Leverage Estimation",
// "description": "Estimate the double leveraged investment return using a constant rate and current market return."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'LeverageMultiplier' of type float64 and assign it a value
	const LeverageMultiplier float64 = 1.35
	// TODO: Declare a Type Inference variable 'marketReturn' (float64) and assign it a value
	marketReturn := 4.75
	// TODO: Calculate Type Inference variable 'leveragedReturn' using formula: (LeverageMultiplier * marketReturn) * 2
	leveragedReturn := (LeverageMultiplier * marketReturn) * 2
	// TODO: Print the value of the 'leveragedReturn' variable
	fmt.Println(leveragedReturn)
}