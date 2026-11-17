// "id": "pr_3_3_b",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "System Integrity",
// "title": "Integrity Watchdog",
// "description": "Monitor system security status until an integrity fault is triggered, using a while-like for loop, constants, and Type Inference Variable Declarations. Declare constants for secure code and strict enforcement, then declare variables for active status, fault counter, and maximum faults. Use a for loop to monitor the system and increment the fault counter until the maximum faults are reached.",
// "answer_options": "for isActive { | faultCounter++ | if faultCounter >= maxFaults { | isActive = false | for isActive = true { | faultCounter+1"
package main

import "fmt"

func main() {
	const secureCode int = 918
	const enforceStrict bool = true
	isActive := true
	faultCounter := 0
	maxFaults := 41
	// TODO: Create a for loop with only a condition to check if 'isActive' is true
	for isActive {
		// TODO: Increment 'faultCounter' by 1 manually inside the loop
		faultCounter++
		// TODO: Create an if statement to check if 'faultCounter' is greater than or equal to 'maxFaults'
		if faultCounter >= maxFaults {
			// TODO: Set 'isActive' to false inside the if statement to stop the loop when maximum faults are reached
			isActive = false
		}
	}
	fmt.Println(secureCode, enforceStrict, faultCounter)
}