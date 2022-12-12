package main

import "fmt"

func main() {
	// new(string) dodeljuje memorijsku lokaciju koja je alocirana za nas
	var firstName *string = new(string)
	fmt.Println(firstName)

	*firstName = "Arthur"
	fmt.Println(firstName)
	fmt.Println(*firstName)

	lastName := "Weasly"
	fmt.Println(lastName)

	// pokazivac ptr
	ptr := &lastName
	fmt.Println(ptr, *ptr)
	lastName = "Booker"
	// memorijska lokacija pokazivaca se ne menja ali vrednost na toj lokaciji da
	fmt.Println(ptr, *ptr)
}
