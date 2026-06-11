package main

import "fmt"


func main(){

	numsA := []int{1,2,3,0,0,0}
	lengthValidNumsA := 3 
	numsB := []int{4,5,6}
	lengthNumsB := len(numsB)

	// Primeiro seguimos para captura dos indices
	maxValidIndexA := lengthValidNumsA - 1
	maxIndexA := lengthValidNumsA + lengthNumsB - 1
	maxIndexB := lengthNumsB - 1

	//Agora começamos uma interação backward com a lista dois
	for maxIndexB >= 0{
		if maxValidIndexA >= 0 && numsA[maxValidIndexA] > numsB[maxIndexB]{ 
			numsA[maxIndexA] = numsA[maxValidIndexA]
			maxValidIndexA --
		} else{
			numsA[maxIndexA] = numsB[maxIndexB]
			maxIndexB --
		}
		maxIndexA--
	}
	fmt.Println(numsA)

}