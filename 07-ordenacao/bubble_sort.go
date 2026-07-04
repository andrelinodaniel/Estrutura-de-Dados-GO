package main

import "fmt"


func bubbleSort1(numeros []int) {
	tamanho := len(numeros)
	for i :=0;i<tamanho-1;i++ {
		for j:= 0; j<tamanho-i-1;j++{
			if numeros[j] < numeros[j+1]{
				numeros[j],numeros[j+1] = numeros[j+1],numeros[j]
			}
		}
	}
}

func main() {
	numeros := []int{9, 4, 7, 1, 3}
	fmt.Println("Antes:", numeros)
	bubbleSort1(numeros)
	fmt.Println("Depois:", numeros)

}
