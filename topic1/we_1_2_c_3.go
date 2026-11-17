// "id": "we_1_2_c_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Fuel Consumption",
// "title": "Total Fuel Burned",
// "description": "Calculate total fuel burned using Type Inference, discard an unused note variable using the blank identifier, and print the results with target limit status."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'burnRate' (float64) and assign it a value
	burnRate := 0.74
	// TODO: Declare a Type Inference variable 'hoursActive' (float64) and assign it a value
	hoursActive := 450.0
	// TODO: Declare and calculate a variable 'totalBurned' as float64 using formula: burnRate * hoursActive
	var totalBurned float64 = burnRate * hoursActive
	// TODO: Declare a Type Inference variable 'activityNote' (string) and assign it a value
	activityNote := "Active run"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'activityNote' variable
	_ = activityNote
	// TODO: Declare and calculate a boolean variable 'isTargetPassed' (bool) using the logical expression (if totalBurned is greater than or equal to 270.0)
	isTargetPassed := totalBurned >= 270.0
	// TODO: Print the calculated totalBurned and isTargetPassed status
	fmt.Println(totalBurned, isTargetPassed)
}