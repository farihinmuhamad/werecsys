// "id": "we_1_3_d_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "System Monitor",
// "title": "Critical System Alert",
// "description": "Evaluate critical system health state using constants, multi-variable boolean declaration, and logical evaluation."
package main

import "fmt"

func main() {
	// TODO: Declare an int variable 'cpuUtilization' and assign it a value
	var cpuUtilization int = 122
	// TODO: Declare 'maxCpuLimit' as a constant int and assignt it a value
	const maxCpuLimit int = 124
	// TODO: Declare boolean type inference Multi Variable 'isOverrideEnabled' and 'isBackupRunning', initialize them with true and false respectively
	isOverrideEnabled, isBackupRunning := true, false
	var isCpuHigh, isOverrideActive, isSystemCritical bool
	// TODO: Determine 'isCpuHigh': True if cpuUtilization is greater than or equal to 'maxCpuLimit'
	isCpuHigh = (cpuUtilization >= maxCpuLimit)
	// TODO: Determine 'isOverrideActive': True if 'isOverrideEnabled' is true OR 'isBackupRunning' is true
	isOverrideActive = (isOverrideEnabled || isBackupRunning)
	// TODO: Determine 'isSystemCritical': Alert is active if 'isCpuHigh' is true AND 'isOverrideActive' is true
	isSystemCritical = isCpuHigh && isOverrideActive
	// TODO: Print the value of 'isSystemCritical'
	fmt.Println("System state critical:", isSystemCritical)
}