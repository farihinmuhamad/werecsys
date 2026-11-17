// "id": "we_1_2_e_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Power System",
// "title": "Power Threshold Check",
// "description": "Check power system input threshold status for voltage and current using type inference and manifest declarations, and evaluate status using logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'voltageInput' using type inference and assign a value
	voltageInput := 4.85
	// TODO: Declare a float64 variable 'currentInput' using type inference and assign a value
	currentInput := 1.58
	// TODO: Declare a manifest float64 variable 'voltageLimit' and assign a value
	var voltageLimit float64 = 4.56
	// TODO: Declare a manifest float64 variable 'currentLimit' and assign a value
	var currentLimit float64 = 1.06
	// TODO: Print the comparison result checking if inputs are greater than limits using logical AND operator
	fmt.Println(voltageInput > voltageLimit && currentInput > currentLimit)
}