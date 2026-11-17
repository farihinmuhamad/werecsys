// "id": "pr_1_2_e",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Smart Home",
// "title": "Sensor Status Check",
// "description": "Check smart home sensor status for temperature and humidity. Declare sensor readings, define thresholds using multi variable declaration, and evaluate status using logical operators.",
// "answer_options": "var sensorRead1 int = 45; var sensorRead2 int = 72; tempLimit, humidityLimit := 38, 76; isWarning := (sensorRead1 > tempLimit) && (sensorRead2 > humidityLimit); sensorRead1, sensorRead2 := 72, 45; tempLimit int, humidityLimit int := 38, 76"
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'sensorRead1' and assign it a value
	var sensorRead1 int = 45
	// TODO: Declare an integer variable 'sensorRead2' and assign it a value
	var sensorRead2 int = 72
	// TODO: Declare two integer variables 'tempLimit' and 'humidityLimit' using multi variable declaration and assign them values
	tempLimit, humidityLimit := 38, 76
	// TODO: Declare a Type Inference variable 'isWarning' (bool) using logical AND operator to check if sensor readings are greater than tempLimit and humidityLimit
	isWarning := (sensorRead1 > tempLimit) && (sensorRead2 > humidityLimit)
	fmt.Println(isWarning)
}