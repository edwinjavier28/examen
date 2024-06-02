package main

import "fmt"

func main33() {

	dicLetras := make(map[int]string)
	dicLetras[1] = "abc"
	dicLetras[2] = "def"
	dicLetras[3] = "ghi"
	dicLetras[4] = "jkl"
	dicLetras[5] = "mno"
	dicLetras[6] = "pqr"
	dicLetras[7] = "stu"
	dicLetras[8] = "vwx"
	dicLetras[9] = "yz0"
	dicLetras[10] = "123"

	for number, letter := range dicLetras {
		fmt.Println(number, letter)
	}

	number := 5
	numero, encontrado := dicLetras[number]
	if encontrado {
		fmt.Println("La letras de", number, numero)
	} else {
		fmt.Println("No se encontraron las letras %s.\n", &number)
	}
	number2 := 7
	numero2, encontrado2 := dicLetras[number2]
	if encontrado2 {
		fmt.Println("La letras de", number2, numero2)
	} else {
		fmt.Println("No se encontraron las letras %s.\n", &number2)
	}
	number3 := 9
	numero3, encontrado3 := dicLetras[number3]
	if encontrado3 {
		fmt.Println("La letras de", number3, numero3)
	} else {
		fmt.Println("No se encontraron las letras %s.\n", &number3)
	}
}
