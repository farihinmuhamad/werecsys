// "id": "we_3_3_c_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Sensor Tracking",
// "title": "Telemetry Parser",
// "description": "Search a series of sensors for a target telemetry packet. Print a parsing message, loop through iterations, skip normal sensors with continue, and stop when the target sensor is found using break. Note: conceptual similarity applied — uses a while-like loop with a condition instead of a standard 3-clause for loop, and type inference variable declarations.",

package main

import "fmt"

func main() {
	parseMode := "Bypass"
	sensorCount, targetSensor := 171, 168
	i := 1
	// TODO: Create a for loop with only a condition to check if 'i' is less than or equal to 'sensorCount'
	for i <= sensorCount {
		// TODO: Create an if statement to check if 'i' is less than 'targetSensor'
		if i < targetSensor {
			i++
			continue
		}
		fmt.Println(parseMode, i)
		// TODO: Create an if statement to check if 'i' is equal to 'targetSensor'
		if i == targetSensor {
			// TODO: Execute 'break' statement to exit the loop when 'i' is equal to 'targetSensor'
			break
		}
		i++
	}
}