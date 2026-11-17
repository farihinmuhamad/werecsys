// "id": "we_1_3_b_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Transaction Limit Check",
// "title": "Order Price Verification",
// "description": "Evaluate if an order price exceeds transaction limits using multi-variable declaration of float64s, constants, blank identifier, and logical operator checks."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 constant 'MaxLimit' and assign it a value
	const MaxLimit float64 = 30000.0
	// TODO: Declare Multi Variable 'ticketPrice' and 'ticketCount' (float64) and assign it a value
	var ticketPrice, ticketCount float64 = 151.0, 167.0
	promoCode := "SAVE30"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'promoCode' variable
	_ = promoCode
	// TODO: Declare a boolean variable 'isVipUser' and assign it a value
	isVipUser := false
	// TODO: Declare a Type Inference variable 'isFlagged' (bool) to check if 'ticketPrice' multiplied by 'ticketCount' is greater than 'MaxLimit' AND 'isVipUser' is false
	isFlagged := (ticketPrice*ticketCount > MaxLimit) && !isVipUser
	// TODO: Print the value of the 'isFlagged' variable
	fmt.Println(isFlagged)
}