// "id": "we_2_3_e_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Boiler Control Diagnostics",
// "title": "Assess Boiler Safety Status",
// "description": "Check boiler system conditions including temperature, system type, and calibration status using a switch statement."
package main

import "fmt"

func main() {
	// TODO: Declare three type inference variables for system type, calibration status, and boiler temperature
	systemType := "Core"
	isSystemCalibrated := true
	boilerTemp := 32.5
	switch {
	// TODO: Case statement to check if boilerTemp is greater than 100.0
	case boilerTemp > 100.0:
		fmt.Println("Emergency Shutdown")
	// TODO: Case statement to check if systemType is Core and isSystemCalibrated is true
	case systemType == "Core" && isSystemCalibrated:
		fmt.Println("Optimal Core Operations")
	case systemType == "Auxiliary" && isSystemCalibrated:
		fmt.Println("Optimal Auxiliary Operations")
	// TODO: Case statement to check if systemType is Backup and isSystemCalibrated is true
	case systemType == "Backup" && isSystemCalibrated:
		fmt.Println("Optimal Backup Operations")
	// TODO: Case statement to check if systemType is a known type but is not calibrated
	case systemType == "Core" || systemType == "Auxiliary" || systemType == "Backup":
		fmt.Println("Recalibration Required")
	// TODO: Default case to handle unrecognized system types or conditions
	default:
		fmt.Println("Status Unrecognized")
	}
}