// "id": "we_1_1_b_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Location Format",
// "title": "Formatting Location Info",
// "description": "Example showing how to combine city and country with a space separator."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'cityName' and assign it a value
	var cityName string = "Paris"
	// TODO: Declare a string variable 'countryName' and assign it a value
	var countryName string = "France"
	// TODO: Concatenate 'cityName' and 'countryName' into a variable 'locationDetails'
	var locationDetails string = cityName + countryName
	// TODO: Print the value of the 'locationDetails' variable
	fmt.Println(locationDetails)
}