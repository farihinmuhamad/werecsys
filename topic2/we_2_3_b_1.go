// "id": "we_2_3_b_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Parcel Transit",
// "title": "Air Cargo Shipping Cost",
// "description": "Calculate cargo shipping costs based on weight and destination using nested if-else statements with float64 variables."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference float64 variable 'cargoWeight' and assign it a value
	cargoWeight := 12.5
	// TODO: Declare a Type Inference float64 variable 'airFreightRate' and assign it a value
	airFreightRate := 2500.5
	// TODO: Declare a Type Inference float64 variable 'seaFreightRate' and assign it a value
	seaFreightRate := 5000.5
	// TODO: Declare a Type Inference boolean variable 'isInternational' and assign it a value
	isInternational := true
	// TODO: Create an if statement to check if 'isInternational' is true
	if isInternational {
		// TODO: Create an if statement to check if 'cargoWeight' is greater than 10.0
		if cargoWeight > 10.0 {
			// TODO: Calculate the total cost for high-weight international air cargo and print the result
			totalCost := (airFreightRate * cargoWeight) * 2.0
			fmt.Println("High air cargo shipping cost:", totalCost)
		} else {
			// TODO: Calculate the total cost for regular international air cargo and print the result
			totalCost := seaFreightRate * cargoWeight
			fmt.Println("Regular air cargo shipping cost:", totalCost)
		}
		// TODO: Create an else statement to handle when 'isInternational' is false
	} else {
		// TODO: Create an if statement to check if 'cargoWeight' is greater than 15.0
		if cargoWeight > 15.0 {
			// TODO: Calculate the total cost for high-weight domestic ocean cargo and print the result
			totalCost := (airFreightRate * cargoWeight) * 2.0
			fmt.Println("High ocean cargo shipping cost:", totalCost)
		} else {
			// TODO: Calculate the total cost for regular domestic ocean cargo and print the result
			totalCost := seaFreightRate * cargoWeight
			fmt.Println("Regular ocean cargo shipping cost:", totalCost)
		}
	}
}
