// "id": "we_3_1_a_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Sensor Accumulation",
// "title": "Cumulative Signal Power",
// "description": "Calculate the cumulative sum of N sequential signal readings using float64 conversion."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'maxSensorReadings' and assign it a value
	maxSensorReadings := 7
	// TODO: Declare a float64 variable 'cumulativeSignal' and assign it a value
	var cumulativeSignal float64 = 0.0
	// TODO: Create a regular for loop to iterate from 1 to 'maxSensorReadings' while incrementing 'i' by 1 at each iteration
	for i := 1; i <= maxSensorReadings; i++ {
		// TODO: Add the current index 'i' as float64 to 'cumulativeSignal'
		cumulativeSignal += float64(i)
	}
	// TODO: Print the value of 'cumulativeSignal' to the console
	fmt.Println(cumulativeSignal)
}