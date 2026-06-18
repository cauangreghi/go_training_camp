package main

import "fmt"

func main(){
	romanNumber := "LCLIII"

	mappedRomanAlgarims := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C':100, 'D': 500, 'M': 1000,
	}


	var current int
	var previous int
	var total int

	for i := len(romanNumber) - 1; i >= 0; i--{
		current = mappedRomanAlgarims[romanNumber[i]]

		if previous > current{
			total -= current
		} else{
			total += current
		}
		previous = current
	}
	fmt.Print(total)
}