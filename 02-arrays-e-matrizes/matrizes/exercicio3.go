package main

import "fmt"

var matriz [3][3]string

func Tabuleiro() {
	for i := 0; i < len(matriz); i++ {
		fmt.Println()
		if i > 0 {
			fmt.Println("------")
		}
		for j := 0; j < len(matriz[0]); j++ {
			if j < 2 {
				fmt.Print(" | ")
			}
		}
	}
}

func main() {
	Tabuleiro()

}
