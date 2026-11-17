// "id": "pr_3_1_c",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Thermal Analysis",
// "title": "Heat Gradient Calculator",
// "description": "Calculate consecutive temperatures along a gradient in a specific zone using a regular for loop, constants, and Regular Manifest Typing Variable Declarations. Declare a constant for step size, a variable for the sensor name and current temperature, then iterate those known values to compute the new temperature.",
// "answer_options": "const StepSize float64 = 3.25 | var sensorName string = "Zone-A" | for i := 1; i <= 4; i++ { | currentTemp += StepSize | for i := 1 to 4 { | currentTemp + StepSize"
package main

import "fmt"

func main() {
	// TODO: Declare a float64 constant 'StepSize' and assign it a value
	const StepSize float64 = 3.25
	// TODO: Declare a string variable 'sensorName' and assign it a value
	var sensorName string = "Zone-A"
	var currentTemp float64 = 24.5
	// TODO: Create a regular for loop to iterate 'i' from 1 to 4 while incrementing 'i' by 1 at each iteration
	for i := 1; i <= 4; i++ {
		// TODO: Add the value of 'StepSize' to 'currentTemp'
		currentTemp += StepSize
	}
	fmt.Println(sensorName, currentTemp)
}