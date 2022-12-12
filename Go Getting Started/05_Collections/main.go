package main

import "fmt"

func main() {
	// ARRAY - fixed sized collection of SIMILAR/SAME data types
	// 0-based index array
	fmt.Println("----------------ARRAYS---------------------")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])

	// implicit initialization syntax
	namesArray := [3]string{"Andrija", "Marko", "Bogdan"}
	fmt.Println(namesArray)
	fmt.Println(namesArray[1])
	fmt.Println("-------------------------------------------")
	// -----------------------------------------------------------------------------------------
	fmt.Println("----------------SLICES---------------------")
	// SLICE - flexible arrays, dinamiclly sized, built on top of arrays
	// you can use array to build a slice off of it

	slice := arr[:] // kreiraj slice of niza arr i hocu slice koji ide od pocetka arr do kraja arr
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr)
	fmt.Println(slice)
	// razlog zasto array i slice imaju isti output je taj sto slice ukazuje na podatke koje niz sadrzi
	// niz je value-type i on ima fizicki podatke 1,2,3 u MEMORIJI
	// slice je neka vrsta pokazivaca koja ukazuje na taj niz koji se nalazi ispod njega
	// sve promene na nizu uticu na slajs i sve promene na slajsu uticu na niz

	// inicijalizacija slice-a, razlikuje se od inicijalizacije niza samo u tome sto nema
	// specificiramo broj podataka koji moze da sadrzi
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)

	mySlice = append(mySlice, 17, 42, 27, 100)
	fmt.Println(mySlice)

	// moguce je kreirati SLICE od SLICE-a
	// kreira slice od slice[1] do slice[5-1] ne ukljucujuci slice[5]
	s2 := mySlice[1:5]
	s3 := mySlice[:2]
	s4 := mySlice[1:2]
	fmt.Println(s2, s3, s4)
	fmt.Println("------------------------------------------")
	// -----------------------------------------------------------------------------------------
	fmt.Println("----------------MAPS----------------------")
	// MAPS - coolection types key-value based
	myMap := map[string]int{"Andrija": 17, "Marko": 12}
	fmt.Println(myMap)
	fmt.Println(myMap["Andrija"])

	myMap["Andrija"] = 9
	fmt.Println(myMap)

	delete(myMap, "Andrija")
	fmt.Println(myMap)
	fmt.Println("------------------------------------------")
	// -----------------------------------------------------------------------------------------
	fmt.Println("----------------STRUCTS----------------------")
	// Jedini tip kolekcije koji nam dozvoljava da koristimo razlicite tipove u njemu
	type User struct {
		ID        int
		FirstName string
		LastName  string
	}

	var user User
	// prilikom inicijalizacije promenljive user tipa User svako polje iz te strukture se
	// inicijalizuje odgovarajucom nil vrednoscu za tip tog polja. int je 0 string je prazan
	// pa dobijemo {0   }
	fmt.Println(user)
	user.ID = 17
	user.FirstName = "Andrija"
	user.LastName = "Vlasacevic"
	fmt.Println(user)
	fmt.Println(user.FirstName)

	goodUserInit := User{12, "Marko", "Vlasacevic"}
	fmt.Println(goodUserInit)

	properUserInit := User{
		ID:        12,
		FirstName: "Milica",
		LastName:  "Grubor",
	}

	fmt.Println(properUserInit)

	// Moguce je kreirati strukturu i na PAKETSKOM NIVOU

}
