package main

import (
	"fmt"
	"os"
	"os/exec"
	"math/rand"
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

func addition() {
	/*
		Be användaren mata in ett tal. Spara värdet i en variabel. Upprepa detta 10 gånger. För
		varje gång lägg till det inmatade värdet till variabelns värde. När det är klart skriv ut-
		Summan av det du matat in är: summan.
	*/
	var number int
	sum := 0
	count := 10

	for i := 0; i < count; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		sum += number
	}
	fmt.Printf("Sum = %d\n", sum)
	
}

func printMinorNumbers(){
	/*
		Skapa ett program där användaren får mata in ett tal. Låt sedan programmet skriva ut
		alla siffor som är mindre än det inmatade talet men större än 0. Lös detta med en
		loop.
	*/
	var number int
	
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	for i := 1; i < number; i++ {
		fmt.Printf("Number: %d\n", i)
	}
}

func rollingDice(){
	/*
		Rolling the dice:
		Kasta två tärningar” och visa resultatet enligt skärmdump ända tills man INTE svarar ”y” eller ”yes” på frågan om igen
	*/
	var answer string
	max := 6
	min := 1
	for  {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Rolling dices... ")
		fmt.Println("The values are: ")
		// generate a random integer between 1 and 6
		fmt.Println(rand.Intn(max-min) + min)
		fmt.Println(rand.Intn(max-min) + min)

		fmt.Print("Do you want to continue rolling dice (y/n): ")
		fmt.Scan(&answer)
		if answer != "yes" && answer != "y" {
			break
		}
	}
}
