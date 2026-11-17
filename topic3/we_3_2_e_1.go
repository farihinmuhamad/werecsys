// "id": "we_3_2_e_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Security Scan",
// "title": "Scan Alert Inspector",
// "description": "Iterate through security log items. Print status message, skip non-suspicious items with continue, and halt scanning on critical alert with break using Type Inference Variable Declarations. Declare string variable for alert keyword, and integer variables for critical ID and max items; then iterate through the items."
package main

import "fmt"

func main() {
	// TODO: Declare a type inference string variable 'alertKeyword' and assign it a value
	alertKeyword := "Anomaly"
	// TODO: Declare two type inference integer variables 'criticalId' and 'maxLogs' and assign them values using multi variable declaration
	criticalId, maxLogs := 134, 156
	// TODO: Create a regular for loop to iterate 'scanCount' from 1 to 'maxLogs' with increment of 1 at each iteration
	for scanCount := 1; scanCount <= maxLogs; scanCount++ {
		// TODO: Create an if statement to check if 'scanCount' is less than 'criticalId' minus 3
		if scanCount < (criticalId - 3) {
			// TODO: Use 'continue' statement to skip non-defective items
			continue
		}
		fmt.Println(alertKeyword, scanCount)
		// TODO: Create an if statement to check if 'scanCount' is equal to 'criticalId'
		if scanCount == criticalId {
			// TODO: Use 'break' statement to halt scanning on critical alert
			break
		}
	}
}