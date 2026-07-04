package main

import "fmt"

func bubbleSort(numeros []int) {
	tamanho := len(numeros)

	for i := 0; i < tamanho-1; i++ {
		fmt.Printf("\nPassada %d\n", i+1)

		for j := 0; j < tamanho-i-1; j++ {
			fmt.Printf("Comparando %d e %d\n", numeros[j], numeros[j+1])

			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
				fmt.Println("Trocou:", numeros)
			} else {
				fmt.Println("Nao trocou:", numeros)
			}
		}
	}
}

func main() {
	numeros := []int{5, 3, 8, 1, 2}

	fmt.Println("Antes:", numeros)
	bubbleSort(numeros)
	fmt.Println("\nDepois:", numeros)
}
