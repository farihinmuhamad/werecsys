// "id": "we_1_1_e_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Physics Sim",
// "title": "Calculating Falling Velocity",
// "description": "Calculate velocity of a falling object using gravity constant and height."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'Gravity' of type float64 and assign it a value
	const Gravity float64 = 9.81
	// TODO: Declare a Type Inference variable 'fallHeight' (float64) and assign it a value
	fallHeight := 45.2
	// TODO: Multiply Gravity and fallHeight, then store it in a new float64 variable 'finalVelocity'
	var finalVelocity float64 = Gravity * fallHeight
	// TODO: Print the value of the 'finalVelocity' variable
	fmt.Println(finalVelocity)
}