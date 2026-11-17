// "id": "we_1_2_c_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Telemetry Track",
// "title": "Distance Calculation",
// "description": "Calculate distance from speed and duration using manifest declarations, discard telemetry notes, and print the results with distance check."
package main

import "fmt"

func main() {
	// TODO: Declare multi-variable 'travelVelocity' and 'duration' as float64 and assign values
	var travelVelocity, duration float64 = 4.25, 75.5
	// TODO: Calculate travelDistance as float64 by multiplying travelVelocity and duration
	var travelDistance float64 = travelVelocity * duration
	// TODO: Declare a manifest string variable 'telemetryMsg' and assign it a value
	var telemetryMsg string = "Calculated via telemetry"
	// TODO: Discard the 'telemetryMsg' variable using the blank identifier
	_ = telemetryMsg
	// TODO: Print the calculated travelDistance and the comparison result directly
	fmt.Println(travelDistance, travelDistance >= 290.0)
}