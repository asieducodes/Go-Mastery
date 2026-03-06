package main

import (
	"fmt"
	//"math"
)

func main() {
	fmt.Println("==================\n   Chapter 03\n==================")
	// Using contants
	const TOTAL float32 = 33
	const price float32 = 12.55
	wage := TOTAL * price
	fmt.Println("You will pay: ", wage)
	balance := 455.00 - wage
	if balance < 0 {
		fmt.Println("Balance: ", 0.00)
	} else {
		fmt.Println("Balance: ", balance)
	}
	// Using Variables
	var name = "String"
	fmt.Println(name)
	// Omittting the variable's assignment
	// Page 102 ;Understanding Pointers
}

/*
//firstValue := 22.54
	//secondValue := math.Round(firstValue)
	//fmt.Println("Value: ", secondValue)
	// Understanding IOTA: keyword used to create series of successive untyped integer constants without needingto assign individual valuesto them.
	const (
		WaterSports = iota
		Soccer
		Chess
		Baseball
	)
	// Defining multiple Constants with a single Statement
	const start, end float32 = 0.00, 12.00
	instock, availableQty := true, 22
	fmt.Println("We have ", availableQty, " left and instock: ", instock)
*/
