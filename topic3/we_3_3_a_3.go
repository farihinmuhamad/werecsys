// "id": "we_3_3_a_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Grid Transmission",
// "title": "Transmission Loss Calculator",
// "description": "Calculate total transmission loss across substation steps using a decrementing for loop. Note: deliberate differences applied — uses a decrementing loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a variable 'stationType' and assign it a value
	var stationType string = "Substation"
	// TODO: Declare a type inference variable 'lossPerStep' (float64) and assign it a value
	lossPerStep := 2.45
	// TODO: Declare a type inference variable 'substationCount' (int) and assign it a value
	substationCount := 102
	// TODO: Declare a type inference variable 'totalGridLoss' (float64) and assign it a value
	totalGridLoss := 0.0
	// TODO: Create a decrementing for loop starting from 'substationCount' down to 1
	for i := substationCount; i >= 1; i-- {
		// TODO: Change totalGridLoss value by adding 'lossPerStep' to 'totalGridLoss'
		totalGridLoss += lossPerStep
	}
	// TODO: Print the values of 'stationType' and 'totalGridLoss'
	fmt.Println(stationType, totalGridLoss)
}