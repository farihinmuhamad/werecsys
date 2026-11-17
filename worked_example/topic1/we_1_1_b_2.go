// "id": "we_1_1_b_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Web Address",
// "title": "Constructing Web Domain",
// "description": "Example showing how to join a domain name and extension using short syntax."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'domainName' using short syntax and assign a value to it
	domainName := "google"
	// TODO: Declare a string variable 'extension' using short syntax and assign a value to it
	extension := "com"
	// TODO: Concatenate 'domainName', a dot, and 'extension' into 'fullAddress'
	fullAddress := domainName + "." + extension
	// TODO: Print the value of the 'fullAddress' variable
	fmt.Println(fullAddress)
}