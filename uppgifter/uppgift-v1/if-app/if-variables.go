package main

import (
	"fmt"
)

func higherLess() {
	var number uint

	fmt.Print("Write a number: ")
	fmt.Scan(&number)

	if number > 10 {
		fmt.Println("Number is higher than 10")
	} else if number < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is equal to 10")
	}
}
func milkPacket() {
	/*
		Be användaren att mata in hur många paket mjölk som finns kvar. Är det mindre än
		10 skriv- Beställ 30 paket. Är det mellan 10 och 20 skriv- Beställ 20 paket. Annars
		skriv-Du behöver inte beställa någon mjölk.
	*/
	var number uint8

	fmt.Print("How many milk packet there is: ")
	fmt.Scan(&number)

	if number >= 10 && number <= 20 {
		fmt.Println("Order 20 milk packets")
	} else if number < 10 {
		fmt.Println("Order 30 milk packets")
	} else {
		fmt.Println("You don't need to order milk")
	}
}

func largestNumber() {
	/*
	   Be användaren mata in två tal.number1 och number2.
	   Lagra det STÖRSTA talet av dom två i en tredje variabel, largest
	*/
	var number1 int
	var number2 int
	var number3 int

	fmt.Print("Enter a first number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter a second number: ")
	fmt.Scan(&number2)

	if number1 > number2 {
		number3 = number1
	} else if number1 < number2 {
		number3 = number2
	} else {
		number3 = number1
	}
	fmt.Printf("Largest number is %d\n", number3)
}

func largestNumber2() {
	var number1 int
	var number2 int
	var number3 int
	var number4 int

	fmt.Print("Enter a first number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter a second number: ")
	fmt.Scan(&number2)
	fmt.Print("Enter a third number: ")
	fmt.Scan(&number3)

	if number1 > number2 && number1 > number3 {
		number4 = number1
	} else if number2 > number1 && number2 > number3 {
		number4 = number2
	} else if number3 > number1 && number3 > number2 {
		number4 = number3
	} else {
		number4 = number1
	}
	fmt.Printf("Largest number is %d\n", number4)
}
func travelPrice() {
	/*
		Be användaren ange vilken kategori den tillhör-vuxen, pensionär, student. Om den är
		pensionär eller student kostar resan 20 kr . Om den är vuxen kostar resan 30 kr . Annars
		skall användaren informeras att den har angett en felaktig kategori.
	*/
	adult := "adult"
	retired := "retired"
	student := "student"
	other := "Wrong category"

	var category string
	fmt.Print("Enter your category: ")
	fmt.Scan(&category)

	if category == adult {
		fmt.Println("At pay: 30 kr")
	} else if category == retired || category == student {
		fmt.Println("At pay: 20 kr")
	} else {
		fmt.Println(other)
	}

}
func birthYear() {
	/*
		Be användaren mata in sitt födelseår. Om det är större eller lika med 1980 men
		mindre än 1990 skriv ut –Du är född på 1980-talet. Om det är mindre än 2000 men
		större än eller lika med 1990 skriv ut- Du är född på 1990-talet. Om det är mindre
		än 1980 eller större än eller lika med 2000, skriv- Du är inte född på 1990 eller
		1980-talen.
	*/
	var birthYear uint16
	fmt.Print("Please enter your birth year: ")
	fmt.Scan(&birthYear)

	if birthYear >= 1980 && birthYear < 1990 {
		fmt.Print("You were born in the 80s")
	} else if birthYear < 2000 && birthYear >= 1990 {
		fmt.Print("You were born in the 90s")
	} else {
		fmt.Print("You weren't born in the 80s nor the 90s")
	}
}
func scandinavia() {
	/*
		Be användaren att mata in namnet på ett land där den bor. Om det är Sverige,
		Danmark, eller Norge skall användare informeras att-Du bor i Skandinavien. Om
		inte meddela användaren att den inte bor i Skandinavien.*/

	var country string

	fmt.Print("Please enter the country you live in: ")
	fmt.Scan(&country)
	if country == "Sweden" || country == "Denmark" || country == "Norway" {
		fmt.Println("You live in Scandinavia")
	} else {
		fmt.Println("You don't live in Scandinavia")
	}

}

func moneyControl() {
	/*
			Be användaren mata in en summa på hur mycket pengar den har. Be sedan
			användaren att ange om den har rabatt.

		    Om summan är mellan 15 och 25 och användaren inte har rabatt skriv – Du kan köpa en liten hamburgare.
		    Om summan är mellan 15 och 25 och användaren har rabatt skriv – Du kan köpa en liten hamburgare och en pommes frites.
		    Om summan är större än 25 och mindre än eller lika med 50 och användaren inte har rabatt skriv – Du kan köpa en stor hamburgare.
		    Om summan är större än 25 och mindre än eller lika med 50 och användaren har rabatt skriv – Du kan köpa en stor hamburgare och pommes frites.
		    Om summan är större än 60 eller om den är mellan 50 och 60 och användaren har rabatt skriv – Du kan köpa ett meal med dryck.
	*/
	var amount int
	discount := 0

	fmt.Print("Please enter the amount of money you actually have: ")
	fmt.Scan(&amount)
	fmt.Print("Enter your discount: ")
	fmt.Scan(&discount)

	if amount >= 15 && amount <= 25 && discount == 0 {
		fmt.Print("You can buy a small hamburger\n ")
	} else if amount >= 15 && amount <= 25 && discount != 0 {
		fmt.Print("You can buy a small hamburger with french frise\n")
	} else if amount > 25 && amount <= 50 && discount == 0 {
		fmt.Print("You can buy a big hamburger\n")
	} else if amount > 25 && amount <= 50 && discount != 0 {
		fmt.Print("You can buy a big hamburger with french frise\n")
	} else if amount > 50 && discount != 0 {
		fmt.Print("You can buy a menu with drink\n")
	}
}
