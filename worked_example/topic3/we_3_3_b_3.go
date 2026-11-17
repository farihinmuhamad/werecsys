// "id": "we_3_3_b_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Access Monitor",
// "title": "Session Countdown",
// "description": "Monitor session timeout countdown status until no attempts remain, using a while-like for loop and short variable declarations. Note: deliberate differences applied — uses a decrementing countdown and short variable declarations without constants.",

package main

import "fmt"

func main() {
	// TODO: Declare a type inference variable 'accessCode' (int) and assign it a value
	accessCode := 754
	// TODO: Declare a type inference variable 'strictMode' (bool) and assign it a value
	strictMode := false
	// TODO: Declare a type inference variable 'accessActive' (bool) and assign it a value
	accessActive := true
	// TODO: Declare a type inference variable 'attemptsLeft' (int) and assign it a value
	attemptsLeft := 141
	// TODO: Declare a type inference variable 'minAttempts' (int) and assign it a value
	minAttempts := 0
	// TODO: Create a for loop with only a condition to check if 'accessActive' is true
	for accessActive {
		// TODO: Decrement 'attemptsLeft' by 1 manually inside the loop
		attemptsLeft--
		// TODO: Create an if statement to check if 'attemptsLeft' is less than or equal to 'minAttempts'
		if attemptsLeft <= minAttempts {
			// TODO: Set 'accessActive' to false inside the if statement to stop the loop when minimum attempts are reached
			accessActive = false
		}
	}
	// TODO: Print the values of 'accessCode', 'strictMode', and 'attemptsLeft'
	fmt.Println(accessCode, strictMode, attemptsLeft)
}