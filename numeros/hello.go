package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
    fmt.Println(quote.Go())
    fmt.Println("10 / 3 =", 10/3.0 )
    fmt.Println("10 % 3 =", 10%3 )
    fmt.Printf("Tipo de 100 * 0.75 = %T",100*0.75)
}
