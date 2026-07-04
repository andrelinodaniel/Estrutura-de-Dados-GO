package main

import "fmt"

// StackNode guarda um caractere (rune) e aponta para o nó abaixo dele na pilha.
type StackNode struct {
	Value rune
	Next  *StackNode
}

// Stack é implementada como uma lista encadeada simples, onde "top" é o topo.
type Stack struct {
	top *StackNode
}

// Push empilha um novo símbolo de abertura no topo.
func (s *Stack) Push(value rune) {
	newNode := &StackNode{Value: value, Next: s.top}
	s.top = newNode
}

// Pop remove e retorna o símbolo do topo. O segundo retorno (bool) indica
// se a operação teve sucesso (pilha não estava vazia).
func (s *Stack) Pop() (rune, bool) {
	if s.top == nil {
		return 0, false
	}
	value := s.top.Value
	s.top = s.top.Next
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Clear esvazia a pilha, permitindo reutilizar a mesma instância.
func (s *Stack) Clear() {
	s.top = nil
}

// IsBalanced verifica se todos os parênteses/colchetes/chaves de expr
// estão corretamente abertos e fechados, na ordem certa (comportamento LIFO).
func (s *Stack) IsBalanced(expr string) bool {
	// mapeia cada símbolo de FECHAMENTO para o símbolo de ABERTURA esperado
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range expr {
		switch char {
		case '(', '[', '{':
			s.Push(char)
		case ')', ']', '}':
			top, ok := s.Pop()
			// falha se a pilha já estava vazia (fechamento sem abertura)
			// ou se o símbolo do topo não é o par esperado
			if !ok || top != pairs[char] {
				return false
			}
		}
		// qualquer outro caractere (letras, números, operadores) é ignorado
	}

	// se sobrou algo na pilha, há abertura sem fechamento
	return s.IsEmpty()
}

func main() {
	stack := &Stack{}

	expr1 := "{[(a+b)*c]}"
	fmt.Printf("%q => %v\n", expr1, stack.IsBalanced(expr1))

	stack.Clear()
	expr2 := "{[(a+b]*c)}"
	fmt.Printf("%q => %v\n", expr2, stack.IsBalanced(expr2))

	stack.Clear()
	expr3 := "(a+b))"
	fmt.Printf("%q => %v\n", expr3, stack.IsBalanced(expr3))
}
