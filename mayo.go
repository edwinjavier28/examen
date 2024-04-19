package main

import "fmt"

func main() {
	// Definimos dos números
	numero1 := 1353
	numero2 := 1354

	// Verificamos si numero1 es mayor que numero2
	if numero1 > numero2 {
		fmt.Printf("%d es mayor que %d\n", numero1, numero2)
	} else {
		fmt.Printf("%d no es mayor que %d\n", numero1, numero2)
	}

	// Verificamos si numero1 es menor que numero2
	if numero1 < numero2 {
		fmt.Printf("%d es menor que %d\n", numero1, numero2)
	} else {
		fmt.Printf("%d no es menor que %d\n", numero1, numero2)
	}
}
