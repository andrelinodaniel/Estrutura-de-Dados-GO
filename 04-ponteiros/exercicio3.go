package main

import "fmt"

func adicionarParValor(m *map[int]string, chave int, valor string){
	(*m)[chave] = valor
}

func removeParValor(m *map[int]string,chave int){
	delete(*m,chave)
}

func main (){
	m := map[int]string{1:"Queijo",2:"Presunto"}
	fmt.Println("MAP Original: ",m)
	adicionarParValor(&m,3,"Carne")
	fmt.Println("MAP Apos adicionar itens: ",m)
}