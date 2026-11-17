// "id": "we_1_1_e_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Gaming Scores",
// "title": "Estimating Final Score",
// "description": "Estimate the final game score using a constant base score and multiplier."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'BaseScore' of type int and assign it a value
	const BaseScore int = 163
	// TODO: Declare a Type Inference variable 'multiplier' (int) and assign it a value
	multiplier := 33
	// TODO: Calculate Type Inference variable 'projectedScore' using formula: (BaseScore * multiplier) + 137
	projectedScore := (BaseScore * multiplier) + 137
	// TODO: Print the value of the 'projectedScore' variable
	fmt.Println(projectedScore)
}