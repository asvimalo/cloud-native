package main

import (
	"fmt"
	"os"
    "os/exec"
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

func minicalculator() {
	/*
			Skapa ett program där användaren

		    Får mata in två tal.
		    Skriv sedan till skärmen summan av de två talen.
		    Skriv sedan en fråga- Vill du fortsätta(J/N)?.
		    Svarar användaren J återupprepas punkt a-c
		    Svarar användaren N avbryts programmet

	*/
	var number1 int
	var number2 int
	//var answer rune
	var answer string

	for {
		cmd := exec.Command("clear")
    	cmd.Stdout = os.Stdout
    	cmd.Run()
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)

		fmt.Printf("Sum = %d\n", number1+number2)

		fmt.Print("Do you want to continue (Y/N)?: ")
		fmt.Scan(&answer)

		if answer == "N" {
			break
		}
	}

}
