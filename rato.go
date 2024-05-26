package main

import "fmt"

type Nodo1 struct {
	letra     string
	siguiente *Nodo
}

type ListaEnlazada struct {
	cabeza *Nodo
}

//func (lista *ListaEnlazada) Insertar(letra string) {
//	nuevoNodo := &Nodo1{letra: letra}
//	if lista.cabeza == nil {
//		lista.cabeza = nuevoNodo
//	} else {
//		actual1 := lista.cabeza
//		for actual.siguiente != nil {
//			actual = actual.siguiente
//		}
//		actual1.siguiente = nuevoNodo
//	}
//}

func (lista *ListaEnlazada) Mostrar() {
	actual := lista.cabeza
	for actual != nil {
		fmt.Println("=>", actual.letra)
		actual = actual.siguiente
	}
}

func main25() {
	lista := ListaEnlazada{}
	lista.Insertar("A")
	lista.Insertar("B")
	lista.Insertar("C")
	lista.Insertar("D")
	lista.Insertar("E")
	lista.Mostrar()
}
