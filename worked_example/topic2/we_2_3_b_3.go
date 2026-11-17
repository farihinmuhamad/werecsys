// "id": "we_2_3_b_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Thermal Safety System",
// "title": "Evaluate Temperature Danger Index",
// "description": "Assess industrial safety temperatures. Note: deliberate differences applied — inverted outer check to check for non-hazardous state first, and inverted inner comparisons with swapped blocks."
package main

import "fmt"

func main() {
	// TODO: Declare three integer variables 'sensorTemp', 'maxLimit', and 'minLimit' using multi variable declaration and assign them values
	var sensorTemp, maxLimit, minLimit int = 15, 25000, 50000
	// TODO: Declare a boolean variable 'isHazardous' and assign it a value
	var isHazardous bool = true
	// TODO: Create an if statement to check if 'isHazardous' is false
	if !isHazardous {
		// TODO: Create an if statement to check if 'sensorTemp' is less than or equal to 15
		if sensorTemp <= 15 {
			// TODO: Calculate the temperature danger index for normal conditions and print the result
			tempDangerIndex := minLimit * sensorTemp
			fmt.Println("Normal temperature level:", tempDangerIndex)
		} else {
			// TODO: Calculate the temperature danger index for extreme conditions and print the result
			tempDangerIndex := (maxLimit * sensorTemp) * 2
			fmt.Println("Extreme temperature level:", tempDangerIndex)
		}
		// TODO: Create an else statement to handle the case when 'isHazardous' is true
	} else {
		// TODO: Create an if statement to check if 'sensorTemp' is less than or equal to 10
		if sensorTemp <= 10 {
			// TODO: Calculate the temperature danger index for hazardous regular conditions and print the result
			tempDangerIndex := minLimit * sensorTemp
			fmt.Println("Hazardous regular level:", tempDangerIndex)
		} else {
			// TODO: Calculate the temperature danger index for hazardous extreme conditions and print the result
			tempDangerIndex := (maxLimit * sensorTemp) * 2
			fmt.Println("Hazardous extreme level:", tempDangerIndex)
		}
	}
}
