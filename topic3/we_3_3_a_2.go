// "id": "we_3_3_a_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Sound Dampening",
// "title": "Acoustic Loss Calculator",
// "description": "Calculate sound dampening across acoustic panels using a while-like loop. Note: conceptual similarity applied — uses a while-like loop with a condition instead of a standard 3-clause for loop, and explicit variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'barrierType' and assign it a value
	const barrierType string = "Foam"
	// TODO: Declare a constant 'absorptionRate' and assign it a value
	const absorptionRate float64 = 3.81
	// TODO: Declare variables 'panelCount', 'totalAbsorption', and 'i' with explicit types and assign them values
	var panelCount int = 98
	// TODO: Declare a variable 'totalAbsorption' and assign it a value
	var totalAbsorption float64 = 0.0
	// TODO: Declare a variable 'i' and assign it a value
	var i int = 1
	// TODO: Create a for loop with only a condition to check if 'i' is less than or equal to 'panelCount'
	for i <= panelCount {
		// TODO: Change totalAbsorption value by adding 'absorptionRate' to 'totalAbsorption'
		totalAbsorption += absorptionRate
		// TODO: Increment 'i' by 1 manually inside the loop
		i++
	}
	// TODO: Print the values of 'barrierType' and 'totalAbsorption'
	fmt.Println(barrierType, totalAbsorption)
}