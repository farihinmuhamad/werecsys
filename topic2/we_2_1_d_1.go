// "id": "we_2_1_d_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Sensor Reading",
// "title": "Check if Sensor Reading is Balanced",
// "description": "Check if a sensor reading value is balanced or unbalanced using modulus and a switch statement. Declare an int16 variable and use a switch to evaluate the remainder when divided by 2."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'sensorReading' and assign it a value
	var sensorReading int = 371
	// TODO: Declare an integer variable 'sensorCheck' to store the result of modulus operation
	var sensorCheck int = sensorReading % 2
	// TODO: Use switch statement to check sensorReading with modulus operation
	switch sensorCheck {
	// TODO: Case statement to check for balanced reading (remainder 0)
	case 0:
		// TODO: Print "Balanced" for case 0
		fmt.Println("Balanced")
	// TODO: Default case to handle unbalanced reading
	default:
		// TODO: Print "Unbalanced" for default case
		fmt.Println("Unbalanced")
	}
}
