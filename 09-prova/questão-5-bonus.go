package main

import "fmt"

type Estatisticas struct {
	Pontos        int
	Saldo_Gols    int
	Gols_Marcados int
	Gols_Sofridos int
}

type Selecao1 struct {
	Nome  string
	Grupo string
	Dados Estatisticas
}



var selecoes = []Selecao1{
	{"Brasil", "A", Estatisticas {9, 7, 10, 3}},
	{"Argentina", "B", Estatisticas {9, 6, 8, 2}},
	{"França", "C", Estatisticas{ 7, 5, 9, 4}},
	{"Espanha", "D",Estatisticas{ 7, 4, 7, 3}},
	{"Alemanha", "E", Estatisticas{6, 3, 6, 3}},
	{"Portugal", "F",Estatisticas{ 6, 2, 5, 3}},
	{"Inglaterra", "G",Estatisticas{ 5, 1, 4, 3}},
	{"Holanda", "H",Estatisticas{ 4, 0, 3, 3}},
}

func GerarRaking(selecoes []Selecao1) {
	tamanho := len(selecoes)
	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-i-1; j++ {
			if selecoes[j].Dados.Pontos < selecoes[j+1].Dados.Pontos {
				selecoes[j], selecoes[j+1] = selecoes[j+1], selecoes[j]
			}
			
		}
	}
}

func SelecionarEquipes(selecoes []Selecao1) []Selecao1 {
	tamanho := len(selecoes)
	var classificados []Selecao1
	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-i-1; j++ {
			if selecoes[j].Dados.Saldo_Gols < selecoes[j+1].Dados.Saldo_Gols {
				selecoes[j], selecoes[j+1] = selecoes[j+1], selecoes[j]
			}

		}
	}
	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-i-1; j++ {
			if selecoes[j].Dados.Gols_Sofridos > selecoes[j+1].Dados.Gols_Sofridos {
				selecoes[j], selecoes[j+1] = selecoes[j+1], selecoes[j]
			}

		}
	}
	for i := 0; i < (tamanho / 2); i++ {
		classificados = append(classificados, selecoes[i])
	}

	return classificados
}

func main() {
	fmt.Println("Selecoes:")
	for _, selecao := range selecoes {
		fmt.Printf("Seleçao: %s ", selecao.Nome)
		fmt.Printf("Grupo: %s ", selecao.Grupo)
		fmt.Printf("Pontos: %d ", selecao.Dados.Pontos)
		fmt.Printf("Saldo de Gols: %d ", selecao.Dados.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ", selecao.Dados.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n", selecao.Dados.Gols_Sofridos)
	}
	fmt.Println("Raking:")
	GerarRaking(selecoes)
	for _, selecao := range selecoes {
		fmt.Printf("Seleçao: %s ", selecao.Nome)
		fmt.Printf("Grupo: %s ", selecao.Grupo)
		fmt.Printf("Pontos: %d ", selecao.Dados.Pontos)
		fmt.Printf("Saldo de Gols: %d ", selecao.Dados.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ", selecao.Dados.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n", selecao.Dados.Gols_Sofridos)
	}
	classificados := SelecionarEquipes(selecoes)
	fmt.Println("Classificados:")
	for _, selecao := range classificados {
		fmt.Printf("Seleçao: %s ", selecao.Dados.Nome)
		fmt.Printf("Grupo: %s ", selecao.Dados.Grupo)
		fmt.Printf("Pontos: %d ", selecao.Dados.Pontos)
		fmt.Printf("Saldo de Gols: %d ", selecao.Dados.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ", selecao.Dados.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n", selecao.Dados.Gols_Sofridos)
	}
}
