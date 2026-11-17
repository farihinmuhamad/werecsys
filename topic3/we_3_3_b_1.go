// "id": "we_3_3_b_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Database Watchdog",
// "title": "Connection Monitor",
// "description": "Monitor database connection status until a threshold of failures is reached, using a while-like for loop, constants, and Type Inference Variable Declarations. Declare constants for watch port and trace mode, then declare variables for monitoring status, fail counter, and maximum failures. Use a for loop to monitor the connection and increment the fail counter until the maximum failures are reached.",

package main

import "fmt"

func main() {
	const watchPort int = 934
	var traceEnabled bool = true
	isMonitoring := true
	failCount := 0
	maxFails := 110
	// TODO: Create a for loop with only a condition to check if 'isMonitoring' is true
	for isMonitoring {
		// TODO: Increment 'failCount' by 1 manually inside the loop
		failCount++
		// TODO: Create an if statement to check if 'failCount' is greater than or equal to 'maxFails'
		if maxFails <= failCount {
			// TODO: Set 'isMonitoring' to false inside the if statement to stop the loop when maximum faults are reached
			isMonitoring = false
		}
	}
	// TODO: Print the values of 'watchPort', 'traceEnabled', and 'failCount' after the loop
	fmt.Println(watchPort, traceEnabled, failCount)
}