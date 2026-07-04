<![CDATA[<div align="center">

# ⭐ Estruturas de Dados em Go

Implementações do zero usando **ponteiros** e **listas encadeadas** — sem bibliotecas externas.

</div>

---

## 🥞 Pilha (Stack) — LIFO

> **Last In, First Out** — O último a entrar é o primeiro a sair.

**Arquivo:** [`pilha/pilha.go`](./pilha/pilha.go)

### Estrutura Interna

```
Topo → [ "Livro C" ] → [ "Livro B" ] → [ "Livro A" ] → nil
```

A pilha é implementada com uma **lista encadeada simples**. Cada `Node` guarda um valor e um ponteiro para o próximo nó. O `Topo` aponta sempre para o elemento mais recente.

### Operações

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Push(val)` | O(1) | Adiciona elemento no topo |
| `Pop()` | O(1) | Remove e retorna o topo |
| `Imprimir()` | O(n) | Exibe a pilha visualmente |

```go
pilha := &Pilha{}
pilha.Push("Livro A")
pilha.Push("Livro B")
pilha.Push("Livro C")
pilha.Pop() // retorna "Livro C"
```

---

### ✅ Pilha de Verificação de Parênteses

**Arquivo:** [`pilha/pilha_parenteses.go`](./pilha/pilha_parenteses.go)

Usa a pilha para verificar se expressões como `{[(a+b)*c]}` estão **balanceadas**.

**Lógica:**
1. Ao encontrar `(`, `[` ou `{` → empilha
2. Ao encontrar `)`, `]` ou `}` → desempilha e verifica se bate com o par esperado
3. Se a pilha estiver vazia ao final → expressão balanceada ✅

```go
stack.IsBalanced("{[(a+b)*c]}") // → true
stack.IsBalanced("{[(a+b]*c)}") // → false
```

---

## 🖨️ Fila (Queue) — FIFO

> **First In, First Out** — O primeiro a entrar é o primeiro a sair.

**Arquivo:** [`fila/fila_impressao.go`](./fila/fila_impressao.go)

### Estrutura Interna

```
front → [Job 1] → [Job 2] → [Job 3] ← rear
```

Implementada com lista encadeada mantendo referências para `front` (próximo a sair) e `rear` (último a entrar). Inclui um contador `totalProcessed` histórico.

### Operações

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Enqueue(val)` | O(1) | Adiciona ao final |
| `Dequeue()` | O(1) | Remove do início; retorna erro se vazia |
| `Size()` | O(1) | Quantidade de itens na fila agora |
| `TotalProcessed()` | O(1) | Total histórico processado |
| `Clear()` | O(1) | Esvazia a fila (preserva histórico) |

```go
queue := &Queue{}
queue.Enqueue(1)
queue.Enqueue(2)
queue.Dequeue()         // retorna 1
queue.TotalProcessed()  // retorna 1
```

---

## 🔗 Lista Duplamente Encadeada

> Cada nó conhece seu anterior **e** seu próximo — permite navegação bidirecional.

**Arquivo:** [`lista-encadeada/lista_duplamente_encadeada.go`](./lista-encadeada/lista_duplamente_encadeada.go)

### Estrutura Interna

```
nil ← [101] ↔ [202] ↔ [303] → nil
       Head                   Tail
```

Cada `Node` tem três campos: `Value`, `Prev` e `Next`. A lista mantém ponteiros para `Head` e `Tail`.

### Operações

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Append(val)` | O(1) | Adiciona ao final |
| `RemoveByValue(val)` | O(n) | Remove a primeira ocorrência do valor |
| `Print()` | O(n) | Exibe `Head → ... → nil` |
| `PrintReverse()` | O(n) | Exibe `Tail → ... → nil` |

```go
list := &DoublyLinkedList{}
list.Append(101)
list.Append(202)
list.Append(303)
list.RemoveByValue(202) // → true
list.PrintReverse()     // [303 <-> 101 -> nil]
```

---

## 🌳 Árvore Binária de Busca (BST)

> Valores menores à esquerda, maiores à direita — busca em O(log n) em média.

**Arquivo:** [`arvore-binaria/arvore_binaria.go`](./arvore-binaria/arvore_binaria.go)

### Estrutura Interna

```
         50
        /  \
      30    70
     /  \  /
   20   40 60
```

### Operações

| Método | Complexidade (média) | Descrição |
|--------|:--------------------:|-----------|
| `Insert(val)` | O(log n) | Insere mantendo a propriedade BST |
| `Contains(val)` | O(log n) | Verifica se o valor existe |
| `Remove(val)` | O(log n) | Remoção nos 3 casos: folha, 1 filho, 2 filhos |
| `InOrder()` | O(n) | Retorna valores em ordem crescente |
| `CountLeaves()` | O(n) | Conta nós folha (sem filhos) |

### Casos de Remoção Implementados

1. **Nó folha** (sem filhos) → simplesmente remove
2. **Um filho** → o filho ocupa o lugar do nó removido
3. **Dois filhos** → substitui pelo **sucessor in-order** (menor da subárvore direita)

```go
tree := &BST{}
for _, c := range []int{50, 30, 70, 20, 40, 60} {
    tree.Insert(c)
}
tree.Contains(40)    // → true
tree.Remove(70)
tree.InOrder()       // → [20 30 40 50 60]
tree.CountLeaves()   // → 3
```

---

## ▶️ Como Executar

```bash
# Pilha
go run 08-estruturas-de-dados/pilha/pilha.go

# Pilha de Parênteses
go run 08-estruturas-de-dados/pilha/pilha_parenteses.go

# Fila
go run 08-estruturas-de-dados/fila/fila_impressao.go

# Lista Duplamente Encadeada
go run 08-estruturas-de-dados/lista-encadeada/lista_duplamente_encadeada.go

# Árvore Binária de Busca
go run 08-estruturas-de-dados/arvore-binaria/arvore_binaria.go
```

---

[← Voltar ao README principal](../README.md)
]]>
