// "id": "we_3_3_d_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Power Management",
// "title": "Battery Level Tracker",
// "description": "Simulate battery discharge tracking in an infinite loop. Decrease charge percentage until the cutoff safety threshold is met, then exit the loop using Type Inference Variable Declarations. Declare variables for battery status, current charge percentage, cutoff percentage, and usage rate. Use a for loop to print the battery status and charge percentage, check if the charge percentage is less than or equal to the cutoff percentage, and break the loop if so. Decrease the charge percentage by the usage rate in each iteration.",

package main

import "fmt"

func main() {
	// TODO: Declare a type inference variable 'batteryStatus' (string) and assign it a value
	batteryStatus := "Discharging"
	// TODO: Declare three type inference variables 'chargePercent' (float64), 'cutoffPercent' (float64), and 'usageRate' (float64) and assign them values
	chargePercent, cutoffPercent, usageRate := 128.50, 10.25, 3.75
	// TODO: Create a for loop without any arguments (infinite loop)
	for {
		fmt.Println(batteryStatus, chargePercent)
		// TODO: Create an if statement to check if 'chargePercent' is less than or equal to 'cutoffPercent'
		if chargePercent <= cutoffPercent {
			// TODO: Use 'break' statement to exit the loop
			break
		}
		// TODO: Decrease 'chargePercent' by 'usageRate' manually inside the loop
		chargePercent -= usageRate
	}
}