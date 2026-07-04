package main

import "fmt"

// TreeNode representa um produto cadastrado, identificado pelo código de barras (Value).
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BST (Binary Search Tree) guarda apenas a raiz. Toda a lógica de organização
// vem da regra: valores menores vão para a esquerda, maiores para a direita.
type BST struct {
	Root *TreeNode
}

// Insert insere um novo código de barras, mantendo a propriedade da BST.
func (t *BST) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

func insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}
	// se value == node.Value, não faz nada (sem duplicados)
	return node
}

// Contains verifica se um código de barras existe na árvore.
func (t *BST) Contains(value int) bool {
	return containsNode(t.Root, value)
}

func containsNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return containsNode(node.Left, value)
	}
	return containsNode(node.Right, value)
}

// Remove tira um produto da árvore, tratando os 3 casos clássicos de remoção em BST.
func (t *BST) Remove(value int) {
	t.Root = removeNode(t.Root, value)
}

func removeNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil // valor não encontrado, nada a fazer
	}

	if value < node.Value {
		node.Left = removeNode(node.Left, value)
	} else if value > node.Value {
		node.Right = removeNode(node.Right, value)
	} else {
		// achamos o nó a remover
		// Caso 1: nó folha (sem filhos)
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// Caso 2: nó com apenas um filho -> o filho toma o lugar do nó
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Caso 3: nó com dois filhos -> substitui pelo sucessor in-order
		// (o menor valor da subárvore direita), depois remove o sucessor de lá
		successor := findMin(node.Right)
		node.Value = successor.Value
		node.Right = removeNode(node.Right, successor.Value)
	}
	return node
}

// findMin desce sempre para a esquerda até achar o menor valor da subárvore.
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InOrder retorna os valores em ordem crescente (esquerda, raiz, direita).
func (t *BST) InOrder() []int {
	var result []int
	inOrderTraverse(t.Root, &result)
	return result
}

func inOrderTraverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inOrderTraverse(node.Left, result)
	*result = append(*result, node.Value)
	inOrderTraverse(node.Right, result)
}

// CountLeaves conta quantos nós são folhas (sem filho esquerdo e sem filho direito).
func (t *BST) CountLeaves() int {
	return countLeavesNode(t.Root)
}

func countLeavesNode(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return countLeavesNode(node.Left) + countLeavesNode(node.Right)
}

func main() {
	tree := &BST{}
	codigos := []int{50, 30, 70, 20, 40, 60}
	for _, c := range codigos {
		tree.Insert(c)
	}

	fmt.Println("Contém 40?", tree.Contains(40))

	tree.Remove(70)

	fmt.Println("InOrder após remoção:", tree.InOrder())
	fmt.Println("CountLeaves após remoção:", tree.CountLeaves())
}
