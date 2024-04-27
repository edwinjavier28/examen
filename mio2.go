package main

import "fmt"

var suma int

func sumarElementos(elementos []int) int {
	suma := 0
	for _, elemento := range elementos {
		suma += elemento
	}
	return suma
}
func invertirLetras(caracteres []string) []string {
	longitud := len(caracteres)
	invertido := make([]string, longitud) // Crear un nuevo slice para almacenar el resultado invertido

	for i := 0; i < longitud; i++ {
		invertido[i] = caracteres[longitud-i-1]
	}
	return invertido
}
func reconocerPalindromo(caracteres []string) bool {
	letras := invertirLetras(caracteres)
	a := recorrerStrings(letras)
	b := recorrerStrings(caracteres)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(letras, "Letras en recorrer palindromo")

	if a == b {
		return true
	} else {
		return false
	}

}
func recorrerStrings(caracteres []string) string {
	var normal string
	for i := 0; i < len(caracteres); i++ {
		normal += caracteres[i]
	}
	return normal

}
func main21() {

	elementos := []int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	caracteres := []string{"r", "a", "t", "o", "n"}
	slice := make([]int, len(elementos))
	slice2 := make([]string, len(caracteres))
	copy(slice, elementos)
	copy(slice2, caracteres)
	fmt.Println(" Slice 1:\n", elementos, "\n", "Slice 2:", "\n", (caracteres))

	suma := sumarElementos(elementos)
	fmt.Printf("La suma de los elementos del Slice 1: %d", suma)

	invertirLetras(caracteres)
	fmt.Println("\nLas letras invertidas del Slice 2:", (caracteres))
	c := reconocerPalindromo(caracteres)
	if c == true {
		fmt.Println(c, "Es palindromo")
	} else {
		fmt.Println(c, "No es palindromo")
	}

}
