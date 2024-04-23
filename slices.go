package main

import "fmt"

func main14() {
	letras := []string{"a", "f", "r", "t", "y", "f", "a", "c", "e", "w", "q"}
	//elementos := []int{9, 13, 8, 1, 51, 30, 11, 29}
	a := prinLetters(letras)
	fmt.Println(a)
	//sumar := sumarslices(elementos)
	//fmt.Println(sumar)
	//comparaSlices(elementos)
}
func sumarslices(slices []int) int {
	var sumatoriaSlices int

	for _, v := range slices {
		sumatoriaSlices += v
	}

	return sumatoriaSlices
}
func comparaSlices(slices []int) {

	for i := 0; i < len(slices)-1; i++ {
		for j := 0; j < len(slices)-i-1; j++ {
			if slices[j] > slices[j+1] {
				slices[j], slices[j+1] = slices[j+1], slices[j]
			}
		}
	}
	fmt.Println(slices)
}

func prinLetters(letras []string) bool {
	for _, v := range letras {
		if v == "f" {
			return true
		}
		fmt.Println(v)
	}
	return false
}
