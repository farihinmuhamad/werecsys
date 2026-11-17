// "id": "pr_3_3_d",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Drone Telemetry",
// "title": "Altitude Monitor",
// "description": "Simulate drone descent tracking in an infinite loop. Decrease altitude until the safety landing threshold is met, then exit the loop using Type Inference Variable Declarations. Declare variables for drone status, current altitude, target altitude, and step decay. Use a for loop to print the drone status and altitude, check if the altitude is less than or equal to the target altitude, and break the loop if so. Decrease the altitude by the step decay value in each iteration.",
// "answer_options": "for { | if altitude <= targetAltitude { | break | altitude -= stepDecay | for droneStatus { | if targetAltitude <= altitude {"
package main

import "fmt"

func main() {
	droneStatus := "Descending"
	altitude := 105.25
	targetAltitude := 12.5
	stepDecay := 8.25
	// TODO: Create a for loop without any arguments (infinite loop)
	for {
		fmt.Println(droneStatus, altitude)
		// TODO: Create an if statement to check if 'altitude' is less than or equal to 'targetAltitude'
		if altitude <= targetAltitude {
			// TODO: Use 'break' statement to exit the loop
			break
		}
		// TODO: Decrease 'altitude' by 'stepDecay' manually inside the loop
		altitude -= stepDecay
	}
}