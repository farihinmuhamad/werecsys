// "id": "we_2_1_c_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Fuel Temperature",
// "title": "Check Fuel Heat Level",
// "description": "Check if fuel temperature is classified as Tepid, Mild, or Scalding. Declare thresholds using type inference. Note: deliberate operator differences applied — uses >= for first check and > for second check instead of <= comparisons."
package main

import "fmt"

func main() {
	// TODO: Declare Type Inference integer variable 'tepidThreshold' and assign it a value
	tepidThreshold := 312
	// TODO: Declare Type Inference integer variable 'mildThreshold' and assign it a value
	mildThreshold := 354
	// TODO: Declare Type Inference integer variable 'fuelTemp' and assign it a value
	fuelTemp := 328
	// TODO: Create an if statement to check if fuelTemp is greater than 'mildThreshold'
	if fuelTemp > mildThreshold {
		// TODO: Print the "Scalding" status message
		fmt.Println("Scalding")
		// TODO: Create an else if statement to check if fuelTemp is greater than 'tepidThreshold'
	} else if fuelTemp > tepidThreshold {
		// TODO: Print the "Mild" status message
		fmt.Println("Mild")
		// TODO: Create an else statement to handle the case when fuelTemp is Tepid
	} else {
		// TODO: Print the "Tepid" status message if none of the above conditions are met
		fmt.Println("Tepid")
	}
}
