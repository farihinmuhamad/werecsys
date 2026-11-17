// "id": "we_1_3_c_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Industrial Boiler",
// "title": "Boiler PSI Range Verification",
// "description": "Verify industrial boiler PSI levels using a new pointer allocation, multi variable threshold declaration, and comparison operators."
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'psiPtr' of type int using the new keyword
	psiPtr := new(int)
	// TODO: Assign a value of 181 to the dereferenced pointer 'psiPtr'
	*psiPtr = 181
	// TODO: Declare integer type inference multi variable 'minPressure' and 'maxPressure' and assign it a value
	minPressure, maxPressure := 169, 191
	// TODO: Declare a Type Inference variable 'isPressureNormal' (bool) to check if dereferenced pointer 'psiPtr' is greater than 'minPressure' and less than 'maxPressure'
	isPressureNormal :=  (*psiPtr < maxPressure) && (*psiPtr > minPressure)
	// TODO: Print the value of 'isPressureNormal'
	fmt.Println(isPressureNormal)
}