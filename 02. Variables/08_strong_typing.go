package main

import "fmt"

func main()	{
	price :=4.50	//float64

	//cups sold in one day
	quantity :=15	//int

	//total income
	//price * quantity (mismatched types float64 and int)
	//total := price * quantity
	total := price * float64(quantity)

	
	fmt.Printf("Total income during a day %.2f", total)

}
