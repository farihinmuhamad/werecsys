// "id": "pr_2_2_e",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Cluster Resource Management",
// "title": "Node Load Status Evaluation",
// "description": "Evaluate a cluster node load status based on CPU and memory usage using a switch statement. The program should categorize the load into 'Critical Overload', 'Warning Alert', 'Moderate Load', or 'Normal Status' based on the given thresholds on cpuUsage and memUsage.",
// "answer_options": "switch {; case cpuUsage > 93 || memUsage > 93:; case cpuUsage > 77 || memUsage > 77:; case cpuUsage > 54 || memUsage > 54:; switch cpuUsage, memUsage {; case cpuUsage || memUsage > 54:"
package main

import "fmt"

func main() {
	var cpuUsage, memUsage int = 82, 94
	// TODO: Use switch statement without an expression to evaluate the resource usage
	switch {
	// TODO: Case statement to check if 'cpuUsage' is greater than 93 OR 'memUsage' is greater than 93
	case cpuUsage > 93 || memUsage > 93:
		fmt.Println("Critical Overload")
	// TODO: Case statement to check if 'cpuUsage' is greater than 77 OR 'memUsage' is greater than 77
	case cpuUsage > 77 || memUsage > 77:
		fmt.Println("Warning Alert")
	// TODO: Case statement to check if 'cpuUsage' is greater than 54 OR 'memUsage' is greater than 54
	case cpuUsage > 54 || memUsage > 54:
		fmt.Println("Moderate Load")
	default:
		fmt.Println("Normal Status")
	}
}