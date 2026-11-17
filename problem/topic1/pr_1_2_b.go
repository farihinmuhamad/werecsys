// "id": "pr_1_2_b",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Geometry",
// "title": "Circle Circumference",
// "description": "Calculate the circumference of a circle using a constant for Pi."
// "answer_options": "const Pi float64 = 3.14159; radius := 5.5; circumference := 2 * Pi * radius; fmt.Println(circumference); var pi int = 3.14159; fmt.Println(Pi * radius * radius)"
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'Pi' of type float64 and assign it a value
	const Pi float64 = 3.14159
	// TODO: Declare a Type Inference variable 'radius' (float64) and assign it a value
	radius := 5.5
	// TODO: Calculate Type Inference variable 'circumference' using formula: 2 * Pi * radius
	circumference := 2 * Pi * radius
	// TODO: Print the value of the 'circumference' variable
	fmt.Println(circumference)
}