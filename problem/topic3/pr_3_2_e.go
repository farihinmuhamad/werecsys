// "id": "pr_3_2_e",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Product Assembly",
// "title": "Assembly Item Inspector",
// "description": "Iterate through assembly items. Print status message, skip non-defective items with continue, and halt assembly on critical defect with break using Type Inference Variable Declarations. Declare string variable for defect keyword, and integer variables for critical bin and max items; then iterate through the items.",
// "answer_options": "for i := 1; i <= maxItems; i++ { | if i < criticalBin { | continue | if i == criticalBin { | for i = 1; i < maxItems; i+2 { | if i <= criticalBin"
package main

import "fmt"

func main() {
	defectKeyword := "Reject"
	criticalBin, maxItems := 47, 57
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'maxItems' with increment of 1 at each iteration
	for i := 1; i <= maxItems; i++ {
		// TODO: Create an if statement to check if 'i' is less than 'criticalBin'
		if i < criticalBin {
			// TODO: Use 'continue' statement to skip non-defective items
			continue
		}
		fmt.Println(defectKeyword, i)
		// TODO: Create an if statement to check if 'i' is equal to 'criticalBin'
		if i == criticalBin {
			break
		}
	}
}