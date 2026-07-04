<![CDATA[<div align="center">

# 📚 Estudos em Go — Estruturas de Dados & Algoritmos

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Status](https://img.shields.io/badge/Status-Em%20Progresso-yellow?style=for-the-badge)]()
[![License](https://img.shields.io/badge/Licença-MIT-green?style=for-the-badge)]()

> Repositório de estudos em **Go (Golang)** cobrindo desde a introdução à linguagem até estruturas de dados clássicas implementadas do zero, com ponteiros e listas encadeadas.

</div>

---

## 🗂️ Estrutura do Repositório

| # | Pasta | Tópico | Conteúdo |
|---|-------|--------|----------|
| 01 | [`01-introducao/`](./01-introducao/) | Introdução ao Go | Hello World, variáveis, entrada/saída |
| 02 | [`02-arrays-e-matrizes/`](./02-arrays-e-matrizes/) | Arrays e Matrizes | Vetores unidimensionais e bidimensionais |
| 03 | [`03-slices/`](./03-slices/) | Slices e Maps | Slices dinâmicos, maps, operações |
| 04 | [`04-ponteiros/`](./04-ponteiros/) | Ponteiros | Referências, passagem por ponteiro |
| 05 | [`05-structs/`](./05-structs/) | Structs | Tipos compostos, métodos, encapsulamento |
| 06 | [`06-recursividade/`](./06-recursividade/) | Recursividade | Funções recursivas, casos base e recursivos |
| 07 | [`07-ordenacao/`](./07-ordenacao/) | Ordenação | Bubble Sort implementado do zero |
| **08** | **[`08-estruturas-de-dados/`](./08-estruturas-de-dados/)** | **Estruturas de Dados** ⭐ | **Pilha, Fila, Lista Encadeada, Árvore BST** |
| 09 | [`09-prova/`](./09-prova/) | Avaliação | Exercícios de prova combinando os tópicos |

---

## ⭐ Destaque: Estruturas de Dados

A pasta [`08-estruturas-de-dados/`](./08-estruturas-de-dados/) é o núcleo deste repositório, com implementações completas usando **ponteiros e listas encadeadas** em Go puro — sem bibliotecas externas.

### Implementações

```
08-estruturas-de-dados/
├── pilha/                    # Stack (LIFO)
│   ├── pilha.go              # Pilha genérica de strings com Push/Pop
│   └── pilha_parenteses.go   # Verificador de expressões balanceadas
├── fila/                     # Queue (FIFO)
│   └── fila_impressao.go     # Fila de impressão com histórico
├── lista-encadeada/          # Linked List
│   └── lista_duplamente_encadeada.go
└── arvore-binaria/           # Binary Search Tree
    └── arvore_binaria.go     # BST com Insert, Remove, InOrder, CountLeaves
```

| Estrutura | Arquivo | Operações |
|-----------|---------|-----------|
| 🥞 **Pilha (Stack)** | `pilha/pilha.go` | `Push`, `Pop`, `Imprimir` |
| ✅ **Pilha de Parênteses** | `pilha/pilha_parenteses.go` | `IsBalanced`, detecção de `()[]{}` |
| 🖨️ **Fila (Queue)** | `fila/fila_impressao.go` | `Enqueue`, `Dequeue`, `TotalProcessed` |
| 🔗 **Lista Duplamente Encadeada** | `lista-encadeada/lista_duplamente_encadeada.go` | `Append`, `RemoveByValue`, `Print`, `PrintReverse` |
| 🌳 **Árvore Binária de Busca** | `arvore-binaria/arvore_binaria.go` | `Insert`, `Contains`, `Remove`, `InOrder`, `CountLeaves` |

---

## 🚀 Como Executar

Você precisa ter o [Go instalado](https://golang.org/dl/) (versão 1.18+).

```bash
# Clone o repositório
git clone https://github.com/<seu-usuario>/GoLang-Estudos.git
cd GoLang-Estudos

# Execute qualquer arquivo diretamente
go run 08-estruturas-de-dados/pilha/pilha.go
go run 08-estruturas-de-dados/arvore-binaria/arvore_binaria.go
go run 07-ordenacao/bubble_sort.go
```

---

## 💡 Conceitos Aplicados

- **Ponteiros** (`*T`, `&variavel`) para manipulação direta de memória
- **Structs com métodos** (receivers) para encapsular comportamento
- **Listas encadeadas** como base para Pilha e Fila
- **Recursividade** em árvores binárias (traversal, contagem, remoção)
- **Interfaces implícitas** do Go para código modular
- **Tratamento de erros** com `error` nativo do Go

---

## 📖 Sobre

Estes estudos fazem parte da disciplina de **Estrutura de Dados** com foco na linguagem Go. Cada implementação foi feita do zero, priorizando clareza, comentários explicativos e boas práticas da linguagem.

---

<div align="center">

Feito com ☕ e muito Go | [@andre](https://github.com)

</div>
]]>
