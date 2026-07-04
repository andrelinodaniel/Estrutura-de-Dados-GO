package main

import "fmt"



type Selecao struct {
	Nome string
	Grupo string
	Pontos int
	Saldo_Gols int
	Gols_Marcados int
	Gols_Sofridos int
}

var selecoes = []Selecao{
   {"Brasil", "A", 9, 7, 10, 3},
   {"Argentina", "B", 9, 6, 8, 2},
   {"França", "C", 7, 5, 9, 4},
   {"Espanha", "D", 7, 4, 7, 3},
   {"Alemanha", "E", 6, 3, 6, 3},
   {"Portugal", "F", 6, 2, 5, 3},
   {"Inglaterra", "G", 5, 1, 4, 3},
   {"Holanda", "H", 4, 0, 3, 3},
}

func GerarRaking(selecoes []Selecao){
	tamanho := len(selecoes)
	for i:=0;i<tamanho -1; i++{
		for j:=0;j<tamanho-i-1;j++{
			if selecoes[j].Pontos < selecoes[j+1].Pontos{
				selecoes[j],selecoes[j+1] = selecoes[j+1],selecoes[j]
			}

		}
	}
}

func SelecionarEquipes(selecoes []Selecao) []Selecao{
	tamanho := len(selecoes)
	var classificados []Selecao 
	for i:=0;i<tamanho -1; i++{
		for j:=0;j<tamanho-i-1;j++{
			if selecoes[j].Saldo_Gols < selecoes[j+1].Saldo_Gols{
				selecoes[j],selecoes[j+1] = selecoes[j+1],selecoes[j]
			}

		}
	}
	for i:=0;i<tamanho -1; i++{
		for j:=0;j<tamanho-i-1;j++{
			if selecoes[j].Gols_Sofridos > selecoes[j+1].Gols_Sofridos{
				selecoes[j],selecoes[j+1] = selecoes[j+1],selecoes[j]
			}

		}
	}
	for i:=0;i<(tamanho/2);i++{
		classificados = append(classificados,selecoes[i])
	}
	

	
	return classificados 
}
func Chaveamento(selecoes  []Selecao) []Selecao,
, []Selecao{
	var chaveamento1  []Selecao
	var chaveamento2 []Selecao

	chaveamento1 = append(chaveamento1,selecoes[0],selecoes[3])
	chaveamento2 =append(chaveamento2,selecoes[1],selecoes[2])
	return chaveamento1,chaveamento2

}
func main() {
	fmt.Println("Selecoes:")
	for _,selecao := range selecoes {
		fmt.Printf("Seleçao: %s ",selecao.Nome)
		fmt.Printf("Grupo: %s ",selecao.Grupo)
		fmt.Printf("Pontos: %d ",selecao.Pontos)
		fmt.Printf("Saldo de Gols: %d ",selecao.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ",selecao.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n",selecao.Gols_Sofridos)
	}
	fmt.Println("Raking:")
	GerarRaking(selecoes)
	for _,selecao := range selecoes {
		fmt.Printf("Seleçao: %s ",selecao.Nome)
		fmt.Printf("Grupo: %s ",selecao.Grupo)
		fmt.Printf("Pontos: %d ",selecao.Pontos)
		fmt.Printf("Saldo de Gols: %d ",selecao.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ",selecao.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n",selecao.Gols_Sofridos)
	}
	classificados := SelecionarEquipes(selecoes)
	fmt.Println("Classificados:")
	for _,selecao := range classificados {
		fmt.Printf("Seleçao: %s ",selecao.Nome)
		fmt.Printf("Grupo: %s ",selecao.Grupo)
		fmt.Printf("Pontos: %d ",selecao.Pontos)
		fmt.Printf("Saldo de Gols: %d ",selecao.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ",selecao.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n",selecao.Gols_Sofridos)

		
	}
	chaveamento1,chaveamento2 = Chaveamento(selecoes)
	fmt.Println("Chaveamento:")
	for _,selecao := range chaveamento1 {
		fmt.Printf("Seleçao: %s ",selecao.Nome)
		fmt.Printf("Grupo: %s ",selecao.Grupo)
		fmt.Printf("Pontos: %d ",selecao.Pontos)
		fmt.Printf("Saldo de Gols: %d ",selecao.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ",selecao.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n",selecao.Gols_Sofridos)

		
	}
	for _,selecao := range chaveamento2 {
		fmt.Printf("Seleçao: %s ",selecao.Nome)
		fmt.Printf("Grupo: %s ",selecao.Grupo)
		fmt.Printf("Pontos: %d ",selecao.Pontos)
		fmt.Printf("Saldo de Gols: %d ",selecao.Saldo_Gols)
		fmt.Printf("Gols Marcados: %d ",selecao.Gols_Marcados)
		fmt.Printf("Gols Sofridos: %d \n",selecao.Gols_Sofridos)

		
	}

}