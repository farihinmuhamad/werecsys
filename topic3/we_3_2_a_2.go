// "id": "we_3_2_a_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Calibration Drift",
// "title": "Drift Offset Tracker",
// "description": "Track cumulative sensor drift by accumulating offset values using a condition-only while-like loop until all calibration steps are complete."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'driftValue' and assign it a value
	var driftValue float64 = 0.88
	// TODO: Declare an integer variable 'calibStep' and assign it a value
	var calibStep int = 0
	// TODO: Declare a float64 variable 'accumulatedDrift' and assign it a value
	var accumulatedDrift float64 = 0.0
	// TODO: Create a while-like for loop that runs while 'calibStep' is less than 26
	for calibStep < 12 {
		// TODO: Update 'accumulatedDrift' by adding 'driftValue' to it
		accumulatedDrift += driftValue
		// TODO: Increment 'calibStep' by 1
		calibStep++
	}
	// TODO: Print the value of 'accumulatedDrift' to the console
	fmt.Println(accumulatedDrift)
}