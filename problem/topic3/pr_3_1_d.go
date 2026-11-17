// "id": "pr_3_1_d",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Submersible Pump Controller",
// "title": "Tank Draining Simulation",
// "description": "Simulate draining a water tank by reducing water level until a low limit is reached, using a while-like for loop and Regular Manifest Typing Variable Declarations. Declare a boolean variable 'pumpActive' to control the loop, and integer variables for water volume and drain limit. Inside the loop, decrement the water volume and check if it has reached or fallen below the drain limit to stop the pump.",
// "answer_options": "for pumpActive {; waterVolume -= 31; if waterVolume <= drainLimit {; pumpActive = false; waterVolume - 31; var pumpActive bool = false"
package main

import "fmt"

func main() {
	var pumpActive bool = true
	var waterVolume, drainLimit int = 97, 27
	// TODO: Create a for loop with only a condition to check if 'pumpActive' is true
	for pumpActive {
		// TODO: Decrement 'waterVolume' by 31 inside the loop
		waterVolume -= 31
		// TODO: Create an if statement to check if 'waterVolume' is less than or equal to 'drainLimit'
		if waterVolume <= drainLimit {
			// TODO: Set 'pumpActive' to false to stop the loop
			pumpActive = false
		}
	}
	fmt.Println(waterVolume)
}