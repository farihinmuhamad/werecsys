// "id": "we_1_3_b_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Lab Experiment",
// "title": "pH Level Safety Verification",
// "description": "Verify chemical sample sterility and safety using constants, float64 variables, blank identifier, and logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 constant 'SafepH' and assign it a value
	const SafepH float64 = 7.45
	// TODO: Declare a float64 variable 'currentpH' and assign it a value
	var currentpH float64 = 7.15
	sampleName := "Sample-A"
	// TODO: Discard 'sampleName' using the underscore blank identifier
	_ = sampleName
	hasContaminants := false
	// TODO: Declare a boolean variable 'isSterile' to check if 'currentpH' is greater than or equal to 'SafepH' AND 'hasContaminants' is false
	var isSterile bool = (currentpH >= SafepH) && !hasContaminants
	fmt.Println(isSterile)
}