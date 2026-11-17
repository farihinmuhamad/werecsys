// "id": "we_3_3_d_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Reservoir Level",
// "title": "Reservoir Level Monitor",
// "description": "Simulate reservoir filling tracking in an infinite loop. Increase water level until the max safety capacity threshold is met, then exit the loop using Type Inference Variable Declarations. Declare variables for flow status, current fluid level, maximum capacity, and inflow rate. Note: deliberate differences applied — uses an incrementing loop and checks greater-than-or-equal break threshold.",

package main

import "fmt"

func main() {
	flowStatus := "Filling"
	fluidLevel := 15.25
	maxCapacity := 95.50
	inflowRate := 4.35
	// TODO: Create a for loop without any arguments (infinite loop)
	for {
		fmt.Println(flowStatus, fluidLevel)
		// TODO: Create an if statement to check if 'fluidLevel' is greater than or equal to 'maxCapacity'
		if fluidLevel >= maxCapacity {
			// TODO: Use 'break' statement to exit the loop
			break
		}
		// TODO: Increase 'fluidLevel' by 'inflowRate' manually inside the loop
		fluidLevel += inflowRate
	}
}