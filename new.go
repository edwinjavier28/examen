package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Nombre string
	Edad   int
}

func main() {
	personas := []Person{
		{"An", 25},
		{"Luis", 30},
		{"Pedro", 25},
		{"Maria", 35},
		{"Oso", 25},
	}

	fmt.Println("Antes de ordenar:", personas)

	// Ordenar el slice por edad y luego por nombre
	sort.Slice(personas, func(i, j int) bool {
		if personas[i].Edad == personas[j].Edad {
			return personas[i].Nombre < personas[j].Nombre
		}
		return personas[i].Edad < personas[j].Edad
	})

	fmt.Println("Después de ordenar por edad y luego por nombre:", personas)
}
