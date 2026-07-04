package main

import "fmt"

type Aluno struct {
	Nome      string
	Matricula int
	Notas     []float64
	Media     float64
	Situacao  string
}

var Alunos []Aluno

func bubbleSort(arr []Aluno) {
	n := int(len(arr))
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].Media < arr[j+1].Media {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
func menu() {

	fmt.Println("1- Cadastrar")
	fmt.Println("2- Excluir")
	fmt.Println("3- Listar")
	fmt.Println("4- Ordenar com BubbleSort")
	fmt.Println("5- Sair")

}
func Adicionar_aluno() {
	var aluno Aluno
	var nota float64
	var soma float64

	fmt.Println("Digite o nome do aluno: ")
	fmt.Scan(&aluno.Nome)
	fmt.Println("Digite a matricula: ")
	fmt.Scan(&aluno.Matricula)
	fmt.Println("Digite as notas: ")
	for i := 0; i < 4; i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scan(&nota)
		aluno.Notas = append(aluno.Notas, nota)
	}
	for i := 0; i < 4; i++ {

		soma += aluno.Notas[i]
	}
	aluno.Media = soma / float64(len(aluno.Notas))
	if aluno.Media >= 7 {
		aluno.Situacao = "Aprovado"
	} else {
		aluno.Situacao = "Reprovado"
	}
	Alunos = append(Alunos, aluno)
	fmt.Println("Aluno Adicionado com sucesso!")
}
func Listar_alunos() {
	for _, aluno := range Alunos {
		fmt.Println("Nome: ", aluno.Nome)
		fmt.Println("Matricula: ", aluno.Matricula)
		fmt.Println("Media: ", aluno.Media)
		fmt.Println("Notas: ")
		for i, nota := range aluno.Notas {
			fmt.Printf("Nota %d: %f \n", i+1, nota)
		}
		fmt.Println("Situação: ", aluno.Situacao)
	}

}
func main() {
	var op int
	for {
		menu()
		fmt.Scan(&op)
		if op == 1 {
			Adicionar_aluno()
		}
		if op == 3 {
			Listar_alunos()
		}
		if op == 4 {
			bubbleSort(Alunos)
		}
		if op == 5 {
			break
		}
		if op == 2{
			
			

		}
	}

}
