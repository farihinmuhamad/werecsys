// "id": "pr_1_3_d",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Logic",
// "title": "Access Control Logic",
// "description": "Combine multiple boolean operators (AND, OR, NOT) to determine access."
// "answer_options": "const minimumAge int = 18; var hasTicket, isVIP bool = true, false; accessA = (age >= minimumAge); accessFinal = accessA && accessB; var minimumAge int = 81; var hasTicket, isVIP bool"
package main

import "fmt"

func main() {
	age := 20
	// TODO: Declare 'minimumAge' as a constant int and assignt it a value
	const minimumAge int = 18
	// TODO: Declare Multi Variable 'hasTicket' and 'isVIP' as a boolean, initialize them with true and false respectively
	var hasTicket, isVIP bool = true, false
	var accessA, accessB, accessFinal bool
	// TODO: Determine 'accessA': True if age is greater than or equal to 'minimumAge'
	accessA = (age >= minimumAge)
	accessB = (hasTicket || isVIP)
	// TODO: Determine 'accessFinal': Access is granted if 'accessA' is true AND 'accessB' is true
	accessFinal = accessA && accessB
	fmt.Println("Access granted:", accessFinal)
}