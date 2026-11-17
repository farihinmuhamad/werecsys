// "id": "we_3_1_c_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Temperature Profiling",
// "title": "Heat Gradient Step Down",
// "description": "Calculate temperatures along a thermal gradient using a decrementing loop and short variable declarations. Note: deliberate differences applied — reverse decrement loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 constant 'gradientStep' and assign it a value
	const gradientStep float64 = 1.30
	// TODO: Declare a string variable 'sensorId' and assign it a value using short declaration
	sensorId := "S-Delta"
	// TODO: Declare a float variable 'probeTemp' and assign it a value using short declaration
	probeTemp := 25.5
	// TODO: Create a regular for loop to iterate 'i' from 5 down to 1 while decrementing 'i' by 1 at each iteration
	for i := 5; i >= 1; i-- {
		// TODO: Add the value of 'gradientStep' to 'probeTemp'
		probeTemp += gradientStep
	}
	// TODO: Print the values of 'sensorId' and 'probeTemp'
	fmt.Println(sensorId, probeTemp)
}