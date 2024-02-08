package main

import (
	"fmt"
)

func loopdecimal() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func loopbetween() {
	/*
		Skapa ett program där användaren får mata in två tal. Låt sedan programmet skriva ut alla
		tal som finns mellan dessa två tal på skärmen. Lös detta med en loop.
	*/

	var num1 int
	var num2 int

	fmt.Print("Enter a number ")
	fmt.Scan(&num1)
	fmt.Print("Enter a second number ")
	fmt.Scan(&num2)

	swap := 0

	if num1 > num2 {
		swap = num1
		num1 = num2
		num2 = swap
	} 
	num1++
	for num1 < num2 {
		fmt.Println(num1)	
		num1++
	}
}
