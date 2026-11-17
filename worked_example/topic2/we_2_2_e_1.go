// "id": "we_2_2_e_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Soil Moisture Analysis",
// "title": "Check Soil Moisture Level",
// "description": "Assess moisture levels from two sensors using float64 parameters and determine the soil moisture category."
package main

import "fmt"

func main() {
	// TODO: Declare two float64 variables 'moistureA' and 'moistureB' using Multi Variable Declaration and assign them values
	var moistureA, moistureB float64 = 82.5, 94.1
	// TODO: Use switch statement without an expression to evaluate the moisture levels
	switch {
	// TODO: Case statement to check if 'moistureA' is greater than 93.0 OR 'moistureB' is greater than 93.0
	case moistureA > 93.0 || moistureB > 93.0:
		// TODO: Print the Saturated Soil status message
		fmt.Println("Saturated Soil")
	// TODO: Case statement to check if 'moistureA' is greater than 77.0 OR 'moistureB' is greater than 77.0
	case moistureA > 77.0 || moistureB > 77.0:
		// TODO: Print the Wet Soil status message
		fmt.Println("Wet Soil")
	// TODO: Case statement to check if 'moistureA' is greater than 54.0 OR 'moistureB' is greater than 54.0
	case moistureA > 54.0 || moistureB > 54.0:
		// TODO: Print the Moist Soil status message
		fmt.Println("Moist Soil")
	default:
		// TODO: Default case to print Dry Soil status message
		fmt.Println("Dry Soil")
	}
}
