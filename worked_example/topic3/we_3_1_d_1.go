// "id": "we_3_1_d_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Vessel Draining Controller",
// "title": "Reactant Pressure Release",
// "description": "Simulate pressure release of a reactor vessel by reducing pressure level until a threshold is reached using a while-like loop and float64 parameters."
package main

import "fmt"

func main() {
	// TODO: Declare a boolean variable 'releaseActive' and assign it a value using short declaration
	releaseActive := true
	// TODO: Declare two float64 variables 'chamberPsi' and 'releaseThreshold' and assign them values using multi variable declaration
	var chamberPsi, releaseThreshold float64 = 84.5, 23.9
	// TODO: Create a for loop with only a condition to check if 'releaseActive' is true
	for releaseActive {
		// TODO: Decrement 'chamberPsi' by 31.2 inside the loop
		chamberPsi -= 12.2
		// TODO: Create an if statement to check if 'releaseThreshold' is greater than or equal to 'chamberPsi'
		if releaseThreshold >= chamberPsi{
			// TODO: Set 'releaseActive' to false to stop the loop
			releaseActive = false
		}
	}
	// TODO: Print the final value of 'chamberPsi' after the loop
	fmt.Println(chamberPsi)
}