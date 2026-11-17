// "id": "pr_3_1_e",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Warehouse Audit",
// "title": "Storage Bin Scanner",
// "description": "Audit storage bins until a damaged item is found, then terminate the audit immediately using break without continue, using Regular Manifest Typing and Multi Variable Declaration. Declare a string constant 'auditStatus' to indicate the status of the audit, and integer variables for maximum bins and the bin number with a defect. Use a for loop to iterate through the bins, printing the audit status and bin number, and break the loop when the max amount of defect bins is found.",
// "answer_options": "const auditStatus string = "Damaged" | for i := 1; i <= maxBins; i++ { | if i == defectBin { | break | if defectBin != i { | for i := 1; i = maxBins; i--"
package main

import "fmt"

func main() {
	// TODO: Declare a string constant 'auditStatus' and assign it a value
	const auditStatus string = "Damaged"
	var maxBins, defectBin int = 53, 43
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'maxBins' while incrementing 'i' by 1 at each iteration
	for i := 1; i <= maxBins; i++ {
		fmt.Println(auditStatus, i)
		// TODO: Create an if statement to check if 'i' is equal to 'defectBin'
		if i == defectBin {
			// TODO: Use 'break' statement to exit the loop when a damaged item is found
			break
		}
	}
}