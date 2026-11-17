// "id": "we_3_2_c_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Sensor Polling",
// "title": "Poll Iteration Loop",
// "description": "Process sensor polling cycles in an infinite loop. Print the poll code and the current cycle, increment the cycle counter, and break when a target cycle is reached using Type Inference Variable Declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'pollCode' and assign it a value
	var pollCode string = "Poll"
	// TODO: Declare an integer variable 'targetCycle' and assign it a value
	var targetCycle int = 112
	// TODO: Declare an integer variable 'currentCycle' and assign it a value
	var currentCycle int = 1
	// TODO: Create a for loop without any arguments (infinite loop)
	for {
		// TODO: Print the 'pollCode' and 'currentCycle' variables in each iteration
		fmt.Println(pollCode, currentCycle)
		// TODO: Create an if statement to check if 'currentCycle' is equal to 'targetCycle'
		if currentCycle == targetCycle {
			// TODO: Use 'break' statement to exit the loop when target cycle is reached
			break
		}
		// TODO: Increment 'currentCycle' by 1 manually inside the loop
		currentCycle++
	}
}