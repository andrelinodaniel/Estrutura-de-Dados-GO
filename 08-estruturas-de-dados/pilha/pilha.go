package main

import "fmt"

// Node representa um nó na nossa lista encadeada da pilha.
type Node struct {
	Valor string
	Prox  *Node
}

// Pilha representa a estrutura de controle da pilha, que mantém o topo.
type Pilha struct {
	Topo *Node
}
 
// Push adiciona um novo elemento no topo da pilha.
func (p *Pilha) Push(val string) {
	novoNode := &Node{Valor: val}
	novoNode.Prox = p.Topo
	p.Topo = novoNode
}

// Pop remove o elemento do topo da pilha e retorna seu valor.
// Se a pilha estiver vazia, retorna uma string vazia.
func (p *Pilha) Pop() string {
	if p.Topo == nil {
		return ""
	}

	valor := p.Topo.Valor
	p.Topo = p.Topo.Prox
	return valor
}

// Imprimir é uma função auxiliar para visualizar a pilha no terminal.
func (p *Pilha) Imprimir() {
	if p.Topo == nil {
		fmt.Println("  [ Pilha vazia ]")
		return
	}

	fmt.Println("  ┌──────────────┐")
	atual := p.Topo
	primeiro := true
	for atual != nil {
		if primeiro {
			fmt.Printf("  │ %-12s │ ← Topo\n", atual.Valor)
			primeiro = false
		} else {
			fmt.Printf("  │ %-12s │\n", atual.Valor)
		}
		atual = atual.Prox
	}
	fmt.Println("  └──────────────┘")
}

func main() {
	fmt.Println("══════════════════════════════════")
	fmt.Println("   PILHA DE STRING COM PONTEIROS  ")
	fmt.Println("══════════════════════════════════")

	// Inicializa a pilha vazia (Topo é nil por padrão)
	pilha := &Pilha{}

	// --- 1. Testando o Push ---
	fmt.Println("\n📥 Empilhando elementos (Push):")
	fmt.Println("  -> Adicionando: \"Livro A\"")
	pilha.Push("Livro A")
	fmt.Println("  -> Adicionando: \"Livro B\"")
	pilha.Push("Livro B")
	fmt.Println("  -> Adicionando: \"Livro C\"")
	pilha.Push("Livro C")

	fmt.Println("\n📋 Estado da pilha:")
	pilha.Imprimir()

	// --- 2. Testando o Pop ---
	fmt.Println("\n📤 Desempilhando elementos (Pop):")
	val1 := pilha.Pop()
	fmt.Printf("  -> Removido: %q\n", val1)

	val2 := pilha.Pop()
	fmt.Printf("  -> Removido: %q\n", val2)

	fmt.Println("\n📋 Estado da pilha após os desempilhamentos:")
	pilha.Imprimir()

	// --- 3. Limpando a pilha ---
	fmt.Println("\n📤 Desempilhando o resto:")
	val3 := pilha.Pop()
	fmt.Printf("  -> Removido: %q\n", val3)

	val4 := pilha.Pop()
	fmt.Printf("  -> Removido (Pilha vazia): %q\n", val4)

	fmt.Println("\n📋 Estado final:")
	pilha.Imprimir()
}
