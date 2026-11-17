// "id": "we_3_1_e_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Quality Inspection",
// "title": "Unit Fault Scanner",
// "description": "Scan production units sequentially and stop immediately when a faulty unit is detected using a regular for loop with break, an integer constant, and a multi variable declaration."
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'inspectionTag' and assign it a value
	const inspectionTag int = 7
	// TODO: Declare integer variables 'maxUnits' and 'failUnit' and assign them values
	var maxUnits, failUnit int = 53, 43
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'maxUnits' while incrementing 'i' by 1 at each iteration
	for i := 1; i <= maxUnits; i++ {
		fmt.Println(inspectionTag, i)
		// TODO: Create an if statement to check if 'i' is equal to 'failUnit'
		if i == failUnit {
			// TODO: Use 'break' statement to exit the loop when a faulty unit is found
			break
		}
	}
}