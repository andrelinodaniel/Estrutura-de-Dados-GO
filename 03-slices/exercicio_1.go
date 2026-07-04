package main

import "fmt"

var Pessoas []string
var m map[string]string

func Adicionar_pessoas() {
	var nome string
	var sexo string
	var opcao int
	fmt.Println("Digite o nome da pessoa: ")
	fmt.Scanln(&nome)
	fmt.Println("Qual o sexo dessa pessoa? (M ou F): ")
	fmt.Scanln(&sexo)
	Pessoas = append(Pessoas, nome, sexo)
	fmt.Println("Pessoa adicionada com sucesso")
	fmt.Println("Deseja cadastrar outra pessoa? (1- Sim / 2- Não): ")
	fmt.Scanln(&opcao)
	if opcao == 1 {
		Adicionar_pessoas()
	}
	if opcao == 2 {
		fmt.Println("Pressione Enter para voltar para o menu...")
		fmt.Scanln()
	}
}
func menu() {
	fmt.Println("1- Inserir pessoa")
	fmt.Println("2- Listar Pessos")
	fmt.Println("3- Quantidade Total de Homens e Mulheres")
	fmt.Println("4- Sair")
}

func main() {

	var opcao int
	for {

		menu()
		fmt.Scanln(&opcao)
		if opcao == 1 {
			Adicionar_pessoas()

		}
		if opcao == 2 {
			fmt.Println(Pessoas)

		}
		if opcao == 3 {
			var Homem int = 0
			var Mulheres int = 0
			for i := 0; i < len(Pessoas); i++ {
				if Pessoas[i] == "M" || Pessoas[i] == "m" {

					Homem++
				}
				if Pessoas[i] == "F" || Pessoas[i] == "f" {

					Mulheres++
				}
			}
			fmt.Println("Homens: ", Homem)
			fmt.Println("Mulheres: ", Mulheres)
			fmt.Println("Pressione Enter para voltar para o menu...")
			fmt.Scanln()
		}
		if opcao == 4 {
			fmt.Println("Saindo...")
			break
		}
	}

}
