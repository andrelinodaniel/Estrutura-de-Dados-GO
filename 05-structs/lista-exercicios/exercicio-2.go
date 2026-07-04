package main

import "fmt"
type Produto struct {
	Nome       string
	Codigo     string
	Preco      float64
	Quantidade int
}

var Produtos = []Produto{
	{Nome: "Caderno", Codigo: "P001", Preco: 12.50, Quantidade: 40},
	{Nome: "Caneta Azul", Codigo: "P002", Preco: 2.30, Quantidade: 120},
	{Nome: "Mochila", Codigo: "P003", Preco: 89.90, Quantidade: 15},
	{Nome: "Régua", Codigo: "P004", Preco: 4.75, Quantidade: 60},
	{Nome: "Estojo", Codigo: "P005", Preco: 18.00, Quantidade: 25},
}

func main() {
	for _, produto := range Produtos {
		fmt.Printf("Produto: %s\n",produto.Codigo)
		fmt.Printf("Nome: %s\n",produto.Nome)
		fmt.Printf("Preço: R$%.2f\n",produto.Preco)
		fmt.Printf("Quantidade: %d\n",produto.Quantidade)
		fmt.Println()

	}
}
