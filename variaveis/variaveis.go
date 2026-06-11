// No GO nós declaramos e depois atribuímos valores as variáveis
// PAdrao do Go é pascalCase
package main

import "fmt"

func main(){
	//declaração
	var nomePessoa string
	//Aqui o GO depois de instanciar uma variável atribui para ela o valor nulo correspondente ao tipo: 0 por padrão, não dará erro de unsigned variablew
	var idadePessoa int 
	//PErmite várias instanciações
	var numero1, numero2, numero3 float64

	//declaração com atribuição
	var nomePessoa2 string = "O cara"

	//atribuição
	nomePessoa = "O cara"

	fmt.Println(nomePessoa, " ", nomePessoa2, "Idade", idadePessoa)

	//Interpolação de string com PrintF o %representa o tipo da variavel
	fmt.Printf("numero 1 = %f | numero 2 = %f | numero3 = %f \n", numero1,numero2,numero3)
	
	// DEcalaração curta (com valor de inferência do tipo) - Marmotinha do GO
	nome := "Cauan Gregui"
	x := 0.45
	fmt.Printf("X = %T e nome = %T\n", x,nome)
	x = 90
	fmt.Printf("X = %T e nome = %T", x,nome)



}