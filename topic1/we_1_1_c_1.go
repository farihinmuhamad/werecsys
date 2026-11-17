// "id": "we_1_1_c_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Fruit Distribution",
// "title": "Distributing Apples",
// "description": "Example showing how to divide two integer variables representing apple counts."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable named 'totalApples' and assign a value to it
	var totalApples int = 103
	// TODO: Declare an integer variable named 'basketCount' and assign a value to it
	var basketCount int = 107
	// TODO: Calculate the division of 'totalApples' by 'basketCount' and store it in a new variable 'applesPerBasket'
	var applesPerBasket int = totalApples / basketCount
	// TODO: Print the value of the 'applesPerBasket' variable
	fmt.Println(applesPerBasket)
}