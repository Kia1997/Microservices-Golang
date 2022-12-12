package main

import "fmt"

// konstanta je sada dostupna u celom paketu ne samo u main funkciji ako je deklarisana
// u konst bloku
const (
	pi      = 3.14
	numOne  = 1
	second  = "second"
	iotaVar = iota      // ima vrednost 3 jer bi prva iota tj. nulta vrednost bila za promenljivu pi
	third   = 3 << iota // 3*2*2*2*2
	forth
)

func main() {
	const c = 4
	fmt.Println(c + 3)
	fmt.Println(c + 1.2)

	const nc int = 3
	fmt.Println(nc + 3)
	fmt.Println(float32(nc) + 1.2)

	fmt.Println(pi, numOne, second)
	fmt.Println(iotaVar, third)
	// iako forth nije deklarisan/inicijalizovan golang uzima vrednost prethodne konstante
	// i nju koristi kao vrednost za konstantu forth tj. 3 << iota samo sto je sada iota 5, a ne 4
	fmt.Println(forth)

}
