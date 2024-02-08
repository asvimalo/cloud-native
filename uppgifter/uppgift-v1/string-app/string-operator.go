package main

import (
	"fmt"
)

func stringOperator() {
	var name string
	//	Be användaren mata sitt namn. Lägg det i en variabel som heter namn
	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name)
	//	Skapa en ny variabel namnUpprepat med namn fem gånger efter varann . Skriv ut den nya variabeln
	for i := 0; i < 5; i++ {
		fmt.Printf("%s", name)
	}

}
