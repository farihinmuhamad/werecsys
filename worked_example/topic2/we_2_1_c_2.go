// "id": "we_2_1_c_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "HVAC System",
// "title": "Check Room Temperature Comfort Level",
// "description": "Check if the room ambient temperature is classified as Mild, Comfortable, or Scorching based on HVAC thresholds. Use manifest variable declarations and if-else if-else branching."
package main

import "fmt"

func main() {
	// TODO: Declare a variable 'ambientTemp' of type int and assign it a value
	var ambientTemp int = 60
	// TODO: Declare a variable 'mildLimit' of type int and assign it a value
	var mildLimit int = 52
	// TODO: Declare a variable 'comfortLimit' of type int and assign it a value
	var comfortLimit int = 69
	// TODO: Create an if statement to check if ambientTemp is Mild
	if ambientTemp <= mildLimit {
		// TODO: Print the mild status message
		fmt.Println(ambientTemp, "Mild")
		// TODO: Create an else if statement to check if ambientTemp is Comfortable
	} else if ambientTemp <= comfortLimit {
		// TODO: Print the comfortable status message
		fmt.Println(ambientTemp, "Comfortable")
		// TODO: Create an else statement to handle the case when ambientTemp is Scorching
	} else {
		// TODO: Print the scorching status message if none of the above conditions are met
		fmt.Println(ambientTemp, "Scorching")
	}
}
