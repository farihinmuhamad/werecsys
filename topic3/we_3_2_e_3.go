// "id": "we_3_2_e_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Warehouse Tracking",
// "title": "Stock Item Inspector",
// "description": "Inspect warehouse stock levels using a decrementing for loop. Note: deliberate differences applied — uses a decrementing loop and checks equality directly instead of using continue/break checks."
package main

import "fmt"

func main() {
	// Declare type inference variable 'stockLabel' (string) and assign it a value
	stockLabel := "Restock"
	// Declare two integer variables 'criticalStock' and 'maxStock' using multi variable declaration and assign them values
	criticalStock, maxStock := 184, 186
	// TODO: Create a decrementing for loop starting from 'maxStock' down to 1
	for i := maxStock; i >= 1; i-- {
		// TODO: Create an if statement to check if 'i' is equal to 'criticalStock'
		if i == criticalStock {
			// TODO: Print 'stockLabel' and the current task index 'i'
			fmt.Println(stockLabel, i)
		}
	}
}