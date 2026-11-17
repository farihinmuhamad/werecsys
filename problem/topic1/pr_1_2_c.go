// "id": "pr_1_2_c",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Fitness Tracking",
// "title": "Daily Calorie Burned",
// "description": "Calculate daily calorie burned from steps using Type Inference, discard an unused note variable using the blank identifier, and print the results with target achievement status.",
// "answer_options": "caloriesPerStep := 0.04; burnedCalories := caloriesPerStep * totalSteps; tempNote := "Tracked via smartwatch"; _ = tempNote; var caloriesPerStep int = 0.04; underscore := tempNote"
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'caloriesPerStep' (float64) and assign it a value
	caloriesPerStep := 0.04
	totalSteps := 8500.0
	// TODO: Calculate Type Inference variable 'burnedCalories' using formula: caloriesPerStep * totalSteps
	burnedCalories := caloriesPerStep * totalSteps
	// TODO: Declare a Type Inference variable 'tempNote' (string) and assign it a value
	tempNote := "Tracked via smartwatch"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'tempNote' variable
	_ = tempNote
	isTargetAchieved := burnedCalories >= 300.0
	fmt.Println(burnedCalories, isTargetAchieved)
}