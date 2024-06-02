package main

import "fmt"

// Definición de un nodo para la lista doblemente enlazada
type Nodo struct {
	caracter  string
	siguiente *Nodo
	anterior  *Nodo
}

// Definición de la lista doblemente enlazada
type ListaDoblementeEnlazada1 struct {
	cabeza *Nodo
	cola   *Nodo
}

// Método para insertar un nuevo nodo al principio de la lista
func (lista *ListaDoblementeEnlazada) InsertarAlPrincipio1(caracter string) {
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

// Método para mostrar la lista en sentido hacia adelante
func (lista *ListaDoblementeEnlazada) MostrarHaciaAdelante1() {
	actual := lista.cabeza
	for actual != nil {
		fmt.Println("=>", actual.caracter)
		actual = actual.siguiente
	}

}

// Método para mostrar la lista en sentido hacia atrás
func (lista *ListaDoblementeEnlazada) MostrarHaciaAtras1() {
	actual := lista.cola
	for actual != nil {
		fmt.Println("=>", actual.caracter)
		actual = actual.anterior
	}

}

func main27() {
	lista := ListaDoblementeEnlazada{}

	// Insertar algunos nodos al principio de la lista
	lista.InsertarAlPrincipio("E")
	lista.InsertarAlPrincipio("D")
	lista.InsertarAlPrincipio("C")
	lista.InsertarAlPrincipio("B")
	lista.InsertarAlPrincipio("A")

	// Mostrar la lista en sentido hacia adelante
	fmt.Println("Lista hacia adelante:")
	lista.MostrarHaciaAdelante1()

	// Mostrar la lista en sentido hacia atrás
	fmt.Println("Lista hacia atrás:")
	lista.MostrarHaciaAtras1()
}
