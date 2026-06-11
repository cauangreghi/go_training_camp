package main

import "fmt"

func main(){

	fmt.Println("teste")
	fmt.Println(`Teste
	para escrever em 
	múltiplas linhas`)
	fmt.Println("Cauan tem ",len("cauan"), " letras")
	fmt.Println("Cauan"[0]) // isso é um byte (int8)
	fmt.Println(string("Cauan"[0])) 
}