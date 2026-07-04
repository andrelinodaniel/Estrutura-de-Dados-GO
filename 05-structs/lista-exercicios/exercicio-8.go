package main

import "fmt"
type Ciclista struct{
	Nome string
	TempoConclusao int
	Premiacao string
}

var Competidores = []Ciclista{
	{Nome: "Mateus", TempoConclusao: 3840},
	{Nome: "Joana", TempoConclusao: 3725},
	{Nome: "Otávio", TempoConclusao: 4010},
	{Nome: "Priscila", TempoConclusao: 3650},
	{Nome: "Sérgio", TempoConclusao: 3905},
	{Nome: "Tatiane", TempoConclusao: 3788},
	{Nome: "Vitor", TempoConclusao: 4120},
}
func bubbleSort(arr []Ciclista) {
	n := int(len(arr))
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].TempoConclusao > arr[j+1].TempoConclusao  {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
func main() {
	fmt.Println("-- PREMIAÇÃO DE CICLISMO ---")
	bubbleSort(Competidores)

	for i := range Competidores{
		if i == 0{
			Competidores[i].Premiacao = "Ouro"
		}else if i == 1{
			Competidores[i].Premiacao = "Prata"
		}else if i == 2{
			Competidores[i].Premiacao = "Bronze"
		}else{
			Competidores[i].Premiacao = "Participante"
		}

		
		
	}

	for _,competidores := range Competidores{
		fmt.Printf("Nome: %s\n",competidores.Nome)
		fmt.Printf("Tempo: %d\n",competidores.TempoConclusao)
		fmt.Printf("Medalha: %s\n",competidores.Premiacao)
		fmt.Println("------------------------------------\n")
	}
}