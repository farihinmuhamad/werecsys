// "id": "we_3_2_e_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Traffic Inspection",
// "title": "Vehicle Speed Inspector",
// "description": "Inspect vehicle speeds using a while-like loop. Note: conceptual similarity applied — uses a while-like loop with a condition instead of a standard 3-clause for loop, and explicit variable declarations."
package main

import "fmt"

func main() {
	var speedLabel string = "Speeding"
	var limitIndex int = 142
	var maxVehicles int = 144
	var i int = 1
	// TODO: Create a for loop with only a condition to check if 'i' is less than or equal to 'maxVehicles'
	for i <= maxVehicles {
		// TODO: Create an if statement to check if 'i' is less than 'limitIndex'
		if i < limitIndex {
			// TODO: Increment 'i' by 1
			i++
			// TODO: Use 'continue' statement to skip normal vehicles
			continue
		}
		// TODO: Print the 'speedLabel' and the current value of 'i'
		fmt.Println(speedLabel, i)
		// TODO: Create an if statement to check if 'i' is equal to 'limitIndex'
		if i == limitIndex {
			// TODO: Use 'break' statement to exit the loop
			break
		}
		// TODO: Increment 'i' by 1
		i++
	}
}