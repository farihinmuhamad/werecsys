// "id": "we_3_1_e_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Inventory Scan",
// "title": "Storage Bin Reverse Scan",
// "description": "Audit storage bins in reverse order and stop immediately when a defective bin is found using a decrementing for loop with break. Note: deliberate differences applied — reverse decrement loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'auditLabel' and assign it a value using short declaration
	auditLabel := "Defective"
	// TODO: Declare integer variables 'totalBins' and 'faultBin' using short multi-variable declaration
	totalBins, faultBin := 23, 11
	// TODO: Create a regular for loop to iterate 'bin' from 'totalBins' down to 1 while decrementing 'bin' by 1 at each iteration
	for bin := totalBins; bin >= 1; bin-- {
		fmt.Println(auditLabel, bin)
		// TODO: Create an if statement to check if 'bin' is equal to 'faultBin'
		if bin == faultBin {
			// TODO: Use 'break' statement to exit the loop when a defective bin is found
			break
		}
	}
}