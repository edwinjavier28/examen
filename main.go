package main

import "fmt"

var b string
var resultado2 int

func main() {
	 Número para el cual queremos imprimir la tabla de multiplicar
	numero1 := 10
	result := factorial(numero1)
	fmt.Println(result)
	numero2 := 9
	result := multiplicar(numero1, numero2)
	fmt.Println(result)
	result2 := Isprimo(numero1)
	fmt.Println(result2)

}

func multiplicar(numero1, numero2 int) string {

	for numero := numero1; numero <= numero2; numero++ {

		 Imprimimos la tabla de multiplicar del número
		fmt.Println("Tabla de multiplicar del", numero, ":")
		for i := 1; i <= 10; i++ {
			resultado := numero * i
			a, _ := fmt.Printf("%d x %d = %d\n", numero, i, resultado)
			b = string(a)
		}
	}

	return string(b)
}

func Isprimo(numero int) string {
	var resultado string

	if numero/numero == 0 {
		resultado = "es primo " + string(numero)
	} else {
		resultado = "no es primo " + string(numero)
	}
	return resultado
}

// dos funciones que busquen Factorial de un numero
func factorial(n int) int {
	// Caso base: si n es 0 o 1, el factorial es 1
	if n == 0 || n == 1 {
		return 1
	}

	// Caso recursivo: n * factorial(n-1)
	return n * factorial(n-1)
}
