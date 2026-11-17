// "id": "we_1_3_c_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Aerodynamics Test",
// "title": "Wind Tunnel Velocity Check",
// "description": "Check wind tunnel velocity boundaries using float64 pointers, dereference comparisons, and separate manifest declarations."
package main

import "fmt"

func main() {
	// TODO: Declare float64 variable 'velocityValue' and assign it a value
	velocityValue := 108.4
	// TODO: Declare a float64 pointer variable 'velocityPtr' and assign it the address of 'velocityValue'
	velocityPtr := &velocityValue
	// TODO: Declare a manifest float64 variable 'minVelocity' and assign a value
	var minVelocity float64 = 84.8
	// TODO: Declare a manifest float64 variable 'maxVelocity' and assign a value
	var maxVelocity float64 = 128.6
	// TODO: Print the comparison result checking if dereferenced 'velocityPtr' is greater than 'minVelocity' AND less than 'maxVelocity'
	fmt.Println(*velocityPtr > minVelocity && *velocityPtr < maxVelocity)
}