// "id": "pr_3_2_a",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Sensor Calibration",
// "title": "Calibration Offset Accumulator",
// "description": "Accumulate offset values over multiple calibration steps using a regular for loop and Regular Manifest Typing Variable Declarations. Declare variables for calibration steps and total offset, then iterate to sum the offsets.",
// "answer_options": "var calibrationSteps int = 26 | var totalOffset float64 = 0.0 | for i := 1; i <= calibrationSteps; i++ { | totalOffset += sensorOffset | for i := 1 to i <= calibrationSteps { | totalOffset + sensorOffset"
package main

import "fmt"

func main() {
	var sensorOffset float64 = 0.75
	// TODO: Declare an integer variable 'calibrationSteps' and assign it a value
	var calibrationSteps int = 26
	// TODO: Declare a float64 variable 'totalOffset' and assign it a value
	var totalOffset float64 = 0.0
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'calibrationSteps' while incrementing i by 1 at each iteration
	for i := 1; i <= calibrationSteps; i++ {
		// TODO: Change the value of 'totalOffset' by adding 'sensorOffset' to it inside the loop
		totalOffset += sensorOffset
	}
	fmt.Println(totalOffset)
}
