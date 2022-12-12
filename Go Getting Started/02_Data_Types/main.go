package main

import "fmt"

func main() {
	// var i int
	// i = 42
	// fmt.Println(i)
	// Nije dobar nacin za deklarisanje i inicijalizovanje promenljive

	var f float32 = 3.14
	fmt.Println(f)

	// implicit initialization syntax
	// golang dodeljuje tip promenljive u zavinosti od dodeljenje vrednosti
	firstName := "Arthur"
	fmt.Println(firstName)
	fmt.Printf("%T", firstName)

	boolVar := true
	fmt.Println(boolVar)

	// tip kompleksnih brojeva
	c := complex(3, 4)
	fmt.Println(c)

	// razdvajanje kompleksnog tipa na realni i imaginarni deo pomocu funkcija real() i imag()
	r, i := real(c), imag(c)
	fmt.Println(r, i)

}
