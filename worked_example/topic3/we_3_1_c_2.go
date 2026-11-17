// "id": "we_3_1_c_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Heat Decay",
// "title": "Cooling Rate Simulator",
// "description": "Simulate cooling cycle counts and temperatures until a stable target threshold is reached using a condition-only loop."
package main

import "fmt"

func main() {
	// TODO: Declare a float variable 'targetTemp' and assign it a value
	var targetTemp float64 = 18.0
	// TODO: Declare a float variable 'currentTemp' and assign it a value
	var currentTemp float64 = 35.5
	// TODO: Declare a float variable 'coolantFactor' and assign it a value
	var coolantFactor float64 = 0.85
	// TODO: Declare an integer variable 'coolCycleCount' and assign it a value
	var coolCycleCount int = 0
	// TODO: Create a condition-only for loop to continue cooling while 'currentTemp' is greater than 'targetTemp'
	for currentTemp > targetTemp {
		// TODO: Update 'currentTemp' by multiplying 'currentTemp' with 'coolantFactor' 
		currentTemp *= coolantFactor
		// TODO: Increment 'coolCycleCount' by 1
		coolCycleCount++
	}
	// TODO: Print the values of 'coolCycleCount' and 'currentTemp'
	fmt.Println(coolCycleCount, currentTemp)
}