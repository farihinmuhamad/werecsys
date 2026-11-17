// "id": "we_3_3_d_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Thermal Cooling",
// "title": "Cooling Cycle Simulator",
// "description": "Simulate temperature cooling in a while-like loop. Decrease temperature until the safe temperature is met, using explicit variable declarations with types. Declare variables for cooling status, current temperature in Celsius, safe temperature, and temperature drop. Note: conceptual similarity applied — uses a while-like loop with a condition instead of a standard 3-clause for loop, and explicit variable declarations.",

package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'coolingStatus' and assign it a value
	var coolingStatus string = "Chilling"
	// TODO: Declare a float64 variable 'tempCelsius' and assign it a value
	var tempCelsius float64 = 78.25
	// TODO: Declare a float64 variable 'safeTemp' and assign it a value
	var safeTemp float64 = 18.75
	// TODO: Declare a float64 variable 'tempDrop' and assign it a value
	var tempDrop float64 = 5.75
	// TODO: Create a for loop with only a condition to check if 'tempCelsius' is greater than 'safeTemp'
	for tempCelsius > safeTemp {
		fmt.Println(coolingStatus, tempCelsius)
		// TODO: Decrease 'tempCelsius' by 'tempDrop' manually inside the loop
		tempCelsius -= tempDrop
	}
}