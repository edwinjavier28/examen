package main

import "fmt"

func main() {

	c := new(Audifonos)
	audifono_L := c.NewAudifonos("rojo", 10, "l", "bose")
	audifono_R := c.NewAudifonos("rojo", 10, "R", "bose")
	Escuchar := c.EscucharAudifonos(audifono_L, audifono_R)
	fmt.Println(Escuchar)
}

type Audifonos struct {
	Color string
	Peso  int
	Lado  string
	Marca string
}

func (c *Audifonos) NewAudifonos(Color string, Peso int, Lado string, Marca string) Audifonos {
	return Audifonos{Color: Color, Peso: Peso, Lado: Lado, Marca: Marca}
}
func (c *Audifonos) EscucharAudifonos(audifono_L, audifono_R Audifonos) int {
	fmt.Println("Escuchando audifono L", audifono_L)
	fmt.Println("Escuchando audifono R", audifono_R)
	return 0
}
