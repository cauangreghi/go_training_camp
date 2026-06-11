package main

import "fmt"

func main(){
	fmt.Println("Entre com um valor")
	// No Scan F eu declaro o tipo da variável que quero receber e aponto (ponteiro na memória) para onde quero atribuí-lo
	var x float64
	//para definir o endereço na memória utiliza-se o &
	fmt.Scanf("%f",&x)

	fmt.Println("Você escolheu o valor", x)
}