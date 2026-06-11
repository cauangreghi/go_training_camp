package main

import "fmt"

func main(){
	fmt.Println("Bom dia! Qual seu nome?")
	var nomePessoa string
	fmt.Scanf("%s", &nomePessoa)
	fmt.Printf("Seu nome é %s, prazer te conhecer", nomePessoa)
}