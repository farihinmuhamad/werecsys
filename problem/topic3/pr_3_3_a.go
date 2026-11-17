// "id": "pr_3_3_a",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Fluid Dynamics",
// "title": "Pressure Drop Calculator",
// "description": "Calculate pressure drop along a pipe at regular intervals using a regular for loop, constants, and Type Inference Variable Declarations. Declare constants for fluid type and flow rate, then declare variables for pipe segments and total pressure drop. Use a for loop to iterate through the pipe segments and calculate the total pressure drop.",
// "answer_options": "pipeSegments := 67 | totalPressureDrop := 0.0 | for i := 1; i <= pipeSegments; i++ { | totalPressureDrop += flowRate | for (i = 1; i <= pipeSegments) do { | totalPressureDrop + flowRate"
package main

import "fmt"

func main() {
	const fluidType string = "Hydraulic Oil"
	const flowRate float64 = 4.15
	// TODO: Declare a type inference variable 'pipeSegments' and assign it a value
	pipeSegments := 67
	// TODO: Declare a type inference variable 'totalPressureDrop' and assign it a value
	totalPressureDrop := 0.0
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'pipeSegments' with increment of 1 at each iteration
	for i := 1; i <= pipeSegments; i++ {
		// TODO: Change totalPressureDrop value by adding 'flowRate' to 'totalPressureDrop'
		totalPressureDrop += flowRate
	}
	fmt.Println(fluidType, totalPressureDrop)
}