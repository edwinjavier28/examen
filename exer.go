package main

import (
	"fmt"
	"sort"
)

// Arreglo de numeros a oredenar
type OrdenAscendente []int

func (a OrdenAscendente) Len() int           { return len(a) }
func (a OrdenAscendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrdenAscendente) Less(i, j int) bool { return a[i] > a[j] }

// Lista doblemente enlazada
type Nodo struct {
	caracter  int
	siguiente *Nodo
	anterior  *Nodo
}

type ListaDoblementeEnlazada struct {
	cabeza *Nodo
	cola   *Nodo
}

func (lista *ListaDoblementeEnlazada) InsertarAlPrincipio(caracter int) {
	nuevoNodo := &Nodo{caracter: caracter}
	if lista.cabeza == nil {
		lista.cabeza = nuevoNodo
		lista.cola = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.cabeza
		lista.cabeza.anterior = nuevoNodo
		lista.cabeza = nuevoNodo
	}
}

func (lista *ListaDoblementeEnlazada) MostrarHaciaAdelante() {
	actual := lista.cabeza
	for actual != nil {
		fmt.Println("=>", actual.caracter)
		actual = actual.siguiente
	}

}

func (lista *ListaDoblementeEnlazada) MostrarHaciaAtras() {
	actual := lista.cola
	for actual != nil {
		fmt.Println("=>", actual.caracter)
		actual = actual.anterior
	}

}
func main() {
	//Primer arreglo
	numeros := []int{7, 2, 9, 4, 3, 8, 1, 6, 5}

	fmt.Println("Numeros:", numeros)

	sort.Ints(numeros)
	fmt.Println("Orden ascendente:", numeros)

	sort.Sort(OrdenAscendente(numeros))
	fmt.Println("Orden descendente:", numeros)
	lista := ListaDoblementeEnlazada{}

	//Lista Doblemente Enlazada
	lista.InsertarAlPrincipio(9)
	lista.InsertarAlPrincipio(8)
	lista.InsertarAlPrincipio(7)
	lista.InsertarAlPrincipio(6)
	lista.InsertarAlPrincipio(5)
	lista.InsertarAlPrincipio(4)
	lista.InsertarAlPrincipio(3)
	lista.InsertarAlPrincipio(2)
	lista.InsertarAlPrincipio(1)

	fmt.Println("Lista hacia adelante:")
	lista.MostrarHaciaAdelante()

	fmt.Println("Lista hacia atrás:")
	lista.MostrarHaciaAtras()

	//Constructor de Objetos
	c := new(Esfero)
	esfero1 := c.NewEsfero("naranja", 3, "Esfero", "Lamy", "Pico")
	esfero2 := c.NewEsfero("rojo", 1, "Pluma", "parker", "Rollerball")
	Ver := c.VerEsfero(esfero1, esfero2)
	fmt.Println(Ver)
}

type Esfero struct {
	Color    string
	Cantidad int
	Tipo     string
	Marca    string
	Modelo   string
}

func (c *Esfero) NewEsfero(Color string, Cantidad int, Tipo string, Marca string, Modelo string) Esfero {
	return Esfero{Color: Color, Cantidad: Cantidad, Tipo: Tipo, Marca: Marca, Modelo: Modelo}
}
func (c *Esfero) VerEsfero(esfero1, esfero2 Esfero) int {
	fmt.Println("Esfero 1", esfero1)
	fmt.Println("Esfero 2", esfero2)
	return 0
}
