package main

import (
	"fmt"
	"math"
)

func main(){
	var numero int
	fmt.Println("Entre com o número")
	fmt.Scanf("%d", &numero)

	resultado := math.Sqrt(float64(numero))
	fmt.Printf("A raiz quadrada de %d é %.2f\n", numero, resultado)
}