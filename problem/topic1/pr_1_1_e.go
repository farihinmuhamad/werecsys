// "id": "pr_1_1_e",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Acoustics",
// "title": "Canyon Depth Estimation",
// "description": "Estimate the depth of a canyon using a constant for the speed of sound and the time elapsed for an echo.",
// "answer_options": "const SoundSpeed float64 = 343.0; echoTime := 4.8; depth := (SoundSpeed * echoTime) / 2.0; fmt.Println(depth); var echoTime int = 4.8; var soundSpeed string = 343.0"
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'SoundSpeed' of type float64 and assign it a value
	const SoundSpeed float64 = 343.0
	// TODO: Declare a Type Inference variable 'echoTime' (float64) and assign it a value
	echoTime := 4.8
	// TODO: Calculate Type Inference variable 'depth' using formula: (SoundSpeed * echoTime) / 2.0
	depth := (SoundSpeed * echoTime) / 2.0
	// TODO: Print the value of the 'depth' variable
	fmt.Println(depth)
}