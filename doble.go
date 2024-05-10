package main

import "fmt"

// Definición de un nodo para la lista doblemente enlazada
type Nodo struct {
	valor     int
	siguiente *Nodo
	anterior  *Nodo
}

// Definición de la lista doblemente enlazada
type ListaDoblementeEnlazada struct {
	cabeza *Nodo
	cola   *Nodo
}

// Método para insertar un nuevo nodo al principio de la lista
func (lista *ListaDoblementeEnlazada) InsertarAlPrincipio(valor int) {
	nuevoNodo := &Nodo{valor: valor}
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
func (lista *ListaDoblementeEnlazada) MostrarHaciaAdelante() {
	actual := lista.cabeza
	for actual != nil {
		fmt.Printf("%d -> ", actual.valor)
		actual = actual.siguiente
	}
	fmt.Println("nil")
}

// Método para mostrar la lista en sentido hacia atrás
func (lista *ListaDoblementeEnlazada) MostrarHaciaAtras() {
	actual := lista.cola
	for actual != nil {
		fmt.Printf("%d -> ", actual.valor)
		actual = actual.anterior
	}
	fmt.Println("nil")
}

func main() {
	lista := ListaDoblementeEnlazada{}

	// Insertar algunos nodos al principio de la lista
	lista.InsertarAlPrincipio(3)
	lista.InsertarAlPrincipio(2)
	lista.InsertarAlPrincipio(1)

	// Mostrar la lista en sentido hacia adelante
	fmt.Println("Lista hacia adelante:")
	lista.MostrarHaciaAdelante()

	// Mostrar la lista en sentido hacia atrás
	fmt.Println("Lista hacia atrás:")
	lista.MostrarHaciaAtras()
}
