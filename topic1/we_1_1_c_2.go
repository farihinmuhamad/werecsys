// "id": "we_1_1_c_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Hourly Earnings",
// "title": "Calculating Total Earnings",
// "description": "Example showing how to multiply two float64 variables using short variable declaration."
package main

import "fmt"

func main() {
	// TODO: Declare float64 variables 'wageRate' and 'billableHours' using short variable declaration and assign them values
	wageRate, billableHours := 85.50, 42.5
	// TODO: Multiply wageRate and billableHours, then store it in a new variable 'totalEarnings'
	totalEarnings := wageRate * billableHours
	// TODO: Print the value of the 'totalEarnings' variable
	fmt.Println(totalEarnings)
}