package main

import "fmt"

type Filme struct {
	Nome      string
	Ano       int
	Genero    string
	Duracao   int
	Avaliacao float64
}

var ListaFilme = []Filme{
	{Nome: "A Vida Invisível", Ano: 2019, Genero: "Drama", Duracao: 139, Avaliacao: 8.4},
	{Nome: "Central do Brasil", Ano: 1998, Genero: "Drama", Duracao: 113, Avaliacao: 9.0},
	{Nome: "Cidade de Deus", Ano: 2002, Genero: "Ação", Duracao: 130, Avaliacao: 8.9},
	{Nome: "Que Horas Ela Volta?", Ano: 2015, Genero: "Drama", Duracao: 112, Avaliacao: 8.1},
	{Nome: "Bacurau", Ano: 2019, Genero: "Suspense", Duracao: 131, Avaliacao: 7.8},
	{Nome: "O Auto da Compadecida", Ano: 2000, Genero: "Comédia", Duracao: 104, Avaliacao: 8.7},
	{Nome: "Ainda Estou Aqui", Ano: 2024, Genero: "Drama", Duracao: 135, Avaliacao: 8.6},
	{Nome: "O Homem do Futuro", Ano: 2011, Genero: "Ficção", Duracao: 106, Avaliacao: 7.2},
}

func main() {
	
}