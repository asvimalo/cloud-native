package main

import (
	"fmt"
)

/*
Skapa en applikation där användaren matar in för och efternamn.

Skriv ut Skriv in ditt förnamn: . Ta emot värdet i en variabel
Låt markören vänta på din inmatning på samma rad.
Gör på samma sätt med efternamnet.
Skriv sedan ut namnen i omvänd ordning.
Se till att resultatet ser ut så här.
*/
func main() {

	var name string
	var surname string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your surname: ")
	fmt.Scan(&surname)

	fmt.Printf("Your name is: %s, %s", surname, name)
}
