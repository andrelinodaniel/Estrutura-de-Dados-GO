package main

import "fmt"

type Aluno struct {
	Nome      string
	Matricula int
	Turma     string 
}

var Alunos = []Aluno{
	{Nome: "Ana Silva", Matricula: 1001, Turma: "1A"},
	{Nome: "Bruno Costa", Matricula: 1002, Turma: "1A"},
	{Nome: "Carla Lima", Matricula: 1003, Turma: "1B"},
	{Nome: "Diego Souza", Matricula: 1004, Turma: "1B"},
	{Nome: "Elisa Rocha", Matricula: 1005, Turma: "1C"},
	{Nome: "", Matricula: 0, Turma: "1C"},
}

func validacaoCadastro(aluno Aluno) bool {
	valido := true

	if aluno.Nome == "" {
		fmt.Println("Erro: nome vazio")
		valido = false
	}

	if aluno.Turma == "" {
		fmt.Println("Erro: Turma vazia")
		valido = false
	}

	if aluno.Matricula <= 0 {
		fmt.Println("Erro: Matricula invalida\n")
		valido = false
	}

	return valido
}

func main() {
	for _, aluno := range Alunos {
		if validacaoCadastro(aluno) {
			fmt.Printf("[%d] %s —  %s\n", aluno.Matricula,aluno.Nome, aluno.Turma)
			
		}
	}
}
