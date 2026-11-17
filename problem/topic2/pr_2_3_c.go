// "id": "pr_2_3_c",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Discount Logic",
// "title": "Customer Membership and Coupon Discounts",
// "description": "Calculate discounts based on membership status, total purchase, and coupon availability. Use nested if statements based on this criteria: If the customer is a member AND their total purchase is greater than 250,000, apply a member discount. If they also have a coupon, apply an additional discount. If the customer is NOT a member AND their total purchase is greater than 500,000, apply a non-member discount. If they have a coupon in this case, apply an additional discount. If none of these conditions are met, print 'No Discount'.",
// "answer_options": "if isMember && totalPurchase > 250000 {; if hasCoupon {; } else if !isMember && totalPurchase > 500000 {; } else {; if isMember + totalPurchase >= 250000; } else if not isMember && totalPurchase <= 500000 {"
package main

import "fmt"

func main() {
	var isMember bool = false
	var totalPurchase int = 575000
	var hasCoupon bool = false
	// TODO: Create an if statement to check if isMember is true and totalPurchase is greater than or equal to 250000
	if isMember && totalPurchase > 250000 {
		// TODO: Create another if statement to check if hasCoupon is true
		if hasCoupon {
			fmt.Println("Member Discount + Coupon Applied")
		} else {
			fmt.Println("Member Discount Applied")
		}
	// TODO: Create an else-if statement to check if isMember is false and totalPurchase is greater than or equal to 500000
	} else if !isMember && totalPurchase > 500000 {
		if hasCoupon {
			fmt.Println("Non-member Discount + Coupon Applied")
		} else {
			fmt.Println("Non-member Discount Applied")
		}
	// TODO: Create an else statement to print "No Discount" if none of the conditions are met
	} else {
		fmt.Println("No Discount")
	}
}