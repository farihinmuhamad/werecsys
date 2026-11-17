// "id": "we_1_1_e_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Geometry Calc",
// "title": "Computing Triangle Metric",
// "description": "Compute a custom triangle metric using height factor constant and base width."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'HeightFactor' of type float64 and assign it a value
	const HeightFactor float64 = 1.75
	// TODO: Declare a float64 variable 'baseWidth' and assign it a value
	var baseWidth float64 = 12.4
	// TODO: Calculate Type Inference variable 'triangleArea' using formula: (HeightFactor * baseWidth) / 2.0
	triangleArea := (HeightFactor * baseWidth) / 2.0
	// TODO: Print the value of the 'triangleArea' variable
	fmt.Println(triangleArea)
}