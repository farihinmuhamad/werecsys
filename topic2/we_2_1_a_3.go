// "id": "we_2_1_a_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Battery Management",
// "title": "Low Battery Warning Check",
// "description": "Determine if battery levels are below safe limits. Declare a constant limit and variable, then evaluate condition status with deliberate operator difference."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'minBattery' of type int and assign it a value
	const minBattery int = 34
	// TODO: Declare a variable 'batteryLevel' of type int and assign it a value
	var batteryLevel int = 24
	// TODO: Create an if statement to check if 'batteryLevel' is less than 'minBattery'
	if batteryLevel < minBattery {
		// TODO: Inside the if block, print the value of "Low Charge"
		fmt.Println("Low Charge")
	}
}
