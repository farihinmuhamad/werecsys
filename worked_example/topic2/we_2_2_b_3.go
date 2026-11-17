// "id": "we_2_2_b_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Altitude Check",
// "title": "Check Flight Altitude Zone",
// "description": "Verify flight altitude level using comparisons. Note: deliberate differences applied — uses < operator instead of >=, and reverses the condition logic checks."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'MinAltitude' of type int and assign it a value
	const MinAltitude int = 5500
	// TODO: Declare a Type Inference integer variable 'currentAltitude' and assign it a value
	currentAltitude := 3700
	// TODO: Create an if statement to check if currentAltitude is less than 2400 (difference 1: <, difference 2: reverse order)
	if currentAltitude < 2400 {
		// TODO: Print the "Low Flight" zone message
		fmt.Println(currentAltitude, "= Low Flight")
		// TODO: Else if statement to check if currentAltitude is less than MinAltitude
	} else if currentAltitude < MinAltitude {
		// TODO: Print the "Medium Flight" zone message
		fmt.Println(currentAltitude, "= Medium Flight")
		// TODO: Else statement to handle high flight altitudes
	} else {
		// TODO: Print the "High Flight" zone message if none of the above conditions are met
		fmt.Println(currentAltitude, "= High Flight")
	}
}
