package main

import "fmt"

func main24() {

	c := new(Audifonos)
	audifono_L := c.NewAudifonos("rojo", 10, "L", "Bose", "F1")
	audifono_R := c.NewAudifonos("rojo", 10, "R", "Bose", "F1")
	Escuchar := c.EscucharAudifonos(audifono_L, audifono_R)
	fmt.Println(Escuchar)
}

type Audifonos struct {
	Color  string
	Peso   int
	Lado   string
	Marca  string
	Modelo string
}

func (c *Audifonos) NewAudifonos(Color string, Peso int, Lado string, Marca string, Modelo string) Audifonos {
	return Audifonos{Color: Color, Peso: Peso, Lado: Lado, Marca: Marca, Modelo: Modelo}
}
func (c *Audifonos) EscucharAudifonos(audifono_L, audifono_R Audifonos) int {
	fmt.Println("Escuchando audifono L", audifono_L)
	fmt.Println("Escuchando audifono R", audifono_R)
	return 0
}
