// "id": "we_3_1_b_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Frame Retransmission",
// "title": "Retry Countdown Log",
// "description": "Simulate packet transmission retry logging in reverse by printing an alert message a certain number of times using a decrementing while-like loop. Note: deliberate differences applied — reverse decrement loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'packetAlert' and assign it a value
	var packetAlert string = "Resending all packet frame"
	// TODO: Declare a integer variable 'remainingTries' and assign it a value
	var remainingTries int = 18
	// TODO: Create a for loop with only a condition to check if 'remainingTries' is greater than 0
	for remainingTries > 0 {
		// TODO: Print the 'packetAlert' to the console
		fmt.Println(packetAlert)
		// TODO: Decrement 'remainingTries' by 1
		remainingTries--
	}
}