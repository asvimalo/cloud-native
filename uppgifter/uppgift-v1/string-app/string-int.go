package main

import (
	"fmt"
)


func stringInt() {
	//Skapa en variabel string name med ditt namn
	name := "Andrés"
    //Skapa en int age med din ålder.
	age := 44
    //Skriv sedan ut Jag heter Kalle (innehållet i name) och är 27(innehållet i age) år.
	fmt.Printf("Jag heter %s och är %d år\n", name, age)
}
