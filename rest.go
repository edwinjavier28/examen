package main

import "fmt"

var resta int

func main() {
	//Se definen los numeros de la resta
	numero1 := 33
	numero2 := 11

	//Se realiza la resta
	resta := numero1 - numero2

	//Se realiza el resultado
	fmt.Printf("%d menos %d da como resultado %d", numero1, numero2, resta)
}
