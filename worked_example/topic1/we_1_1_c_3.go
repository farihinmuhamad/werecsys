// "id": "we_1_1_c_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Currency Exchange",
// "title": "Converting Dollars to Euros",
// "description": "Example showing how to convert USD to EUR by multiplying two float64 variables."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'dollars' and assign it a value
	var dollars float64 = 125.50
	// TODO: Declare a type inference float64 variable 'exchangeRate' and assign it a value
	exchangeRate := 1.18
	// TODO: Multiply dollars and exchangeRate, then store it in a new float64 variable 'euros'
	var euros float64 = dollars * exchangeRate
	// TODO: Print the value of the 'euros' variable
	fmt.Println(euros)
}