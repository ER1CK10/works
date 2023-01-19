package main

import "fmt"

const ebulicaoK float64 = 373

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Printf("o ponto de ebulicao em Kelvin é: %g e em Celsius é: %g", tempK, tempC)

}
