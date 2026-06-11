package main

import (
	"fmt"
)

func main(){
	inicio := 0
	listNames := []string{"Aaron", "Abner", "Ademir", "Adriano", "Alan", "Alberto", "Alessandra", "Alexandre", "Alice", "Aline", "Amanda", "Ana", "Anderson", "André", "Andréia", "Ângela", "Antônio", "Arthur", "Beatriz", "Bianca", "Breno", "Bruna", "Bruno", "Caio", "Camila", "Carlos", "Carolina", "Caroline", "César", "Cíntia", "Cláudia", "Cláudio", "Cristian", "Cristina", "Daniel", "Daniela", "Daniele", "Danilo", "Davi", "David", "Débora", "Diego", "Diogo", "Douglas", "Eduarda", "Eduardo", "Elaine", "Eliane", "Elias", "Elisa", "Elisabete", "Eloá", "Emanuel", "Emerson", "Emily", "Eric", "Erica", "Erika", "Esther", "Evelyn", "Fabiana", "Fabiano", "Fábio", "Felipe", "Fernanda", "Fernando", "Filipe", "Flávia", "Flávio", "Franciele", "Francine", "Gabriel", "Gabriela", "Gabrielle", "Geovana", "Giovana", "Giovani", "Gisela", "Gisele", "Giuliano", "Glaucia", "Guilherme", "Gustavo", "Heitor", "Helena", "Henrique", "Hugo", "Ian", "Igor", "Isaac", "Isabel", "Isabela", "Isabella", "Isadora", "Ítalo", "Ivan", "Janaína", "Jaqueline", "Jean", "Jefferson"}

	var meio int
	var counter int
	nomeProcurado := "Jefferson"

	final := len(listNames) - 1
	for inicio <= final	{
		meio = (inicio + final) / 2
		nomeAtual := listNames[meio]

		if nomeAtual == nomeProcurado{
			fmt.Println("NOME ENCONTRADO ", nomeProcurado)
			break
		} else if nomeAtual > nomeProcurado{
			final = meio - 1
			counter++
			continue
		}	else if nomeAtual < nomeProcurado{
			inicio = meio + 1
		}
		counter++
	}
	fmt.Printf("Realizamos %d interações", counter)
}