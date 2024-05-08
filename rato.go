package main

import "fmt"

type Nodo struct {
	letra     string
	siguiente *Nodo
}

type ListaEnlazada struct {
	cabeza *Nodo
}

func (lista *ListaEnlazada) Insertar(letra string) {
	nuevoNodo := &Nodo{letra: letra}
	if lista.cabeza == nil {
		lista.cabeza = nuevoNodo
	} else {
		actual := lista.cabeza
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
	}
}

func (lista *ListaEnlazada) Mostrar() {
	actual := lista.cabeza
	for actual != nil {
		fmt.Println("=>", actual.letra)
		actual = actual.siguiente
	}
}

func main() {
	lista := ListaEnlazada{}
	lista.Insertar("A")
	lista.Insertar("B")
	lista.Insertar("C")
	lista.Insertar("D")
	lista.Insertar("E")
	lista.Mostrar()
}
