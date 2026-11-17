// "id": "we_3_1_e_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Bin Audit Tracker",
// "title": "Storage Scan Terminator",
// "description": "Scan storage bins using a while-like condition-only loop with a manual index, stopping when the scan index exceeds the defect threshold."
package main

import "fmt"

func main() {
	// TODO: Declare a boolean variable 'binScanActive' and assign it a value
	var binScanActive bool = true
	// TODO: Declare two integer variables 'scanIndex' and 'defectThreshold' and assign them values using multi variable declaration
	var scanIndex, defectThreshold int = 1, 43
	// TODO: Create a while-like for loop that runs while 'binScanActive' is true
	for binScanActive {
		// TODO: Print the current scan index
		fmt.Println("Scanning bin:", scanIndex)
		// TODO: Check if scanIndex has reached the defect threshold
		if scanIndex >= defectThreshold {
			// TODO: Set 'binScanActive' to false to stop the loop
			binScanActive = false
		}
		// TODO: Increment the scan index by 1
		scanIndex++
	}
}