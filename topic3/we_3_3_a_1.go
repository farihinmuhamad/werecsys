// "id": "we_3_3_a_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Thermal Dynamics",
// "title": "Heat Loss Calculator",
// "description": "Calculate total heat loss across multiple wall layers at regular intervals using a regular for loop, constants, and Type Inference Variable Declarations. Declare constants for insulation type and loss rate, then declare variables for wall layers and total heat loss. Use a for loop to iterate through the wall layers and calculate the total heat loss."
package main

import "fmt"

func main() {
	const insulationType string = "Fiberglass"
	const lossRate float64 = 5.25
	// TODO: Declare a variable 'wallLayers' and assign it a value
	var wallLayers int = 140
	// TODO: Declare a variable 'totalHeatLoss' and assign it a value
	var totalHeatLoss float64 = 0.0
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'wallLayers' with increment of 1 at each iteration
	for i := 1; i <= wallLayers; i++ {
		// TODO: Change totalHeatLoss value by adding 'lossRate' to 'totalHeatLoss'
		totalHeatLoss += lossRate
	}
	fmt.Println(insulationType, totalHeatLoss)
}