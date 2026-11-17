// "id": "we_3_1_d_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Pump Fluid Control",
// "title": "Tank Draining Break Control",
// "description": "Simulate draining a water tank using a while-like loop that breaks when water volume drops below a limit. Note: deliberate differences applied — short variable declarations and break statement."
package main

import "fmt"

func main() {
	// TODO: Declare a boolean variable 'pumpRunning' and assign it a value using short declaration
	pumpRunning := true
	// TODO: Declare two integer variables 'reservoirVolume' and 'thresholdLimit' and assign them values using multi variable declaration
	var reservoirVolume, thresholdLimit int = 162, 29
	// TODO: Create a for loop with only a condition to check if 'pumpRunning' is true
	for pumpRunning {
		// TODO: Decrement 'reservoirVolume' by 21 inside the loop
		reservoirVolume -= 21
		// TODO: Create an if statement to check if 'reservoirVolume' is less than or equal to 'thresholdLimit'
		if reservoirVolume <= thresholdLimit {
			// TODO: Use a break statement to stop the loop
			break
		}
	}
	// TODO: Print the final value of 'reservoirVolume' after the loop
	fmt.Println(reservoirVolume)
}