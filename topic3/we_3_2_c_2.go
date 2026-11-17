// "id": "we_3_2_c_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Reset Sequence",
// "title": "Reset Phase Loop",
// "description": "Execute device reset phases using a condition-only while-like loop. Note: conceptual similarity applied — uses a while-like loop with a condition instead of an infinite loop."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'resetMarker' and assign it a value
	var resetMarker string = "Phase"
	// TODO: Declare an integer variable 'targetPhase' and assign it a value
	var targetPhase int = 126
	// TODO: Declare an integer variable 'currentPhase' and assign it a value
	var currentPhase int = 1
	// TODO: Create a for loop with a condition to check if 'currentPhase' is less than or equal to 'targetPhase'
	for currentPhase <= targetPhase {
		fmt.Println(resetMarker, currentPhase)
		// TODO: Increment 'currentPhase' by 1 manually inside the loop
		currentPhase++
	}
}