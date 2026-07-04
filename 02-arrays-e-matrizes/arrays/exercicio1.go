// 8. Dado o slice []int{12, 45, 7, 23, 56, 89, 34, 67, 11, 90}.
// Ordene o slice em ordem crescente. Use busca binaria para verificar se o valor 56
// esta presente e imprima o indice encontrado.
package main

import (
	"fmt"
	"slices"
)

func main() {
	numeros := []int{12, 45, 7, 23, 56, 89, 34, 67, 11, 90}

	valor := 56


	if encontrado {
		fmt.Printf("Valor %d encontrado no indice %d\n", valor, indice)
	} else {
		fmt.Println("Valor nao encontrado!")
	}
}
