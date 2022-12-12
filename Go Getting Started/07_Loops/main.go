package main

func main() {
	// All loops are for loops
	var i int
	// sve dok je i manje od 5 ovo ce se vrteti, ako se i ne menja onda je beskonacno
	// LOOP 1
	for i < 5 {
		i++
		if i == 1 {
			continue
		}
		println(i)
		if i == 2 {
			break
		}
	}

	// LOOP 2
	// impicitna sintaksa (i:=0) za i ne menja vrednost i van for petlje, vec lokalno
	println("Loop 2")
	for i := 0; i < 5; i++ {
		println(i)
	} // poslednje i ce biti 5

	println(i) // i se vraca na 2

	// LOOP 3 - INFINITE LOOP
	println("Infinite loop")
	for {
		if i == 5 {
			break
		}
		i++
		println(i)
	}

	// LOOP 4 - ARRAY/SLICE/MAP LOOP
	println("Slice loop 1")
	slice := []string{"Andrija", "Marko", "Bogdan"}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	println("Slice loop 2")
	// ovo moze za sve kolekcije
	for index, value := range slice {
		println(index, value)
	}

	wellKnownPorts := map[string]int{"http": 80, "ssh": 23, "https": 443}
	for key, value := range wellKnownPorts {
		println(key, value)
	}

	// moze i bez vrednost
	for key := range wellKnownPorts {
		println(key)
	}

	// moze i samo vrednosti
	for _, values := range wellKnownPorts {
		println(values)
	}

	println("______________________________________________________________________")

	panic("Something bad happened")
}
