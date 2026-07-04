<div align="center">

<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original-wordmark.svg" alt="Go Logo" width="120"/>

# Estrutura de Dados em Go

**Estudos completos de estruturas de dados e algoritmos implementados do zero em Go puro**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![GitHub repo size](https://img.shields.io/github/repo-size/andrelinodaniel/Estrutura-de-Dados-GO?style=flat-square&color=00ADD8)](https://github.com/andrelinodaniel/Estrutura-de-Dados-GO)
[![Status](https://img.shields.io/badge/status-em%20progresso-f0a500?style=flat-square)]()
[![License](https://img.shields.io/badge/licença-MIT-2ea44f?style=flat-square)]()

<br/>

> Cada estrutura foi implementada **do zero**, sem bibliotecas externas,  
> usando **ponteiros**, **listas encadeadas** e **recursividade** com Go idiomático.

</div>

---

## 📖 Sobre o Repositório

Este repositório reúne minha jornada de estudos em **Go (Golang)** durante a disciplina de **Estrutura de Dados**. O conteúdo vai desde os fundamentos da linguagem até implementações completas das principais estruturas clássicas da Ciência da Computação.

Cada pasta representa um módulo de estudo independente, com exemplos comentados e exercícios resolvidos.

---

## 🗂️ Módulos de Estudo

| Módulo | Tópico | O que você vai encontrar |
|:------:|--------|--------------------------|
| [`01`](./01-introducao/) | 🟦 **Introdução ao Go** | Hello World, variáveis, tipos, entrada/saída com `fmt` |
| [`02`](./02-arrays-e-matrizes/) | 📦 **Arrays e Matrizes** | Vetores unidimensionais, matrizes 2D, iteração |
| [`03`](./03-slices/) | 🔀 **Slices e Maps** | Slices dinâmicos, `append`, `make`, mapas chave-valor |
| [`04`](./04-ponteiros/) | 📌 **Ponteiros** | `*T`, `&var`, passagem por referência, manipulação de memória |
| [`05`](./05-structs/) | 🏗️ **Structs** | Tipos compostos, métodos com receivers, encapsulamento |
| [`06`](./06-recursividade/) | 🔁 **Recursividade** | Casos base, chamadas recursivas, pilha de chamadas |
| [`07`](./07-ordenacao/) | 🫧 **Ordenação** | Bubble Sort implementado e otimizado do zero |
| [`08`](./08-estruturas-de-dados/) | ⭐ **Estruturas de Dados** | Pilha, Fila, Lista Encadeada, Árvore BST |
| [`09`](./09-prova/) | 📝 **Avaliação** | Questões de prova integrando todos os conceitos |

---

## ⭐ Estruturas de Dados — Núcleo do Repositório

> Todas implementadas com **listas encadeadas e ponteiros** — sem slices, sem `container/list`.

<br/>

### 🥞 Pilha — Stack (LIFO)

```
  Topo →  ┌──────────────┐
          │ Livro C      │ ← push/pop aqui
          ├──────────────┤
          │ Livro B      │
          ├──────────────┤
          │ Livro A      │
          └──────────────┘
```

**Arquivos:** [`pilha/pilha.go`](./08-estruturas-de-dados/pilha/pilha.go) · [`pilha/pilha_parenteses.go`](./08-estruturas-de-dados/pilha/pilha_parenteses.go)

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Push(val)` | **O(1)** | Empilha um elemento no topo |
| `Pop()` | **O(1)** | Desempilha e retorna o topo |
| `IsBalanced(expr)` | **O(n)** | Verifica parênteses `()[]{}` balanceados |

---

### 🖨️ Fila — Queue (FIFO)

```
  Enqueue →  [ Job 3 ] [ Job 2 ] [ Job 1 ]  → Dequeue
                rear                 front
```

**Arquivo:** [`fila/fila_impressao.go`](./08-estruturas-de-dados/fila/fila_impressao.go)

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Enqueue(val)` | **O(1)** | Adiciona ao final da fila |
| `Dequeue()` | **O(1)** | Remove do início; retorna `error` se vazia |
| `TotalProcessed()` | **O(1)** | Contador histórico (nunca decresce) |

---

### 🔗 Lista Duplamente Encadeada

```
  nil ←  [ 101 ] ↔ [ 202 ] ↔ [ 303 ]  → nil
          Head                  Tail
```

**Arquivo:** [`lista-encadeada/lista_duplamente_encadeada.go`](./08-estruturas-de-dados/lista-encadeada/lista_duplamente_encadeada.go)

| Método | Complexidade | Descrição |
|--------|:-----------:|-----------|
| `Append(val)` | **O(1)** | Insere no final via ponteiro `Tail` |
| `RemoveByValue(val)` | **O(n)** | Remove primeira ocorrência, ajustando `Prev`/`Next` |
| `PrintReverse()` | **O(n)** | Percorre do `Tail` ao `Head` |

---

### 🌳 Árvore Binária de Busca — BST

```
           50
          /  \
        30    70
       /  \  /
      20  40 60
```

**Arquivo:** [`arvore-binaria/arvore_binaria.go`](./08-estruturas-de-dados/arvore-binaria/arvore_binaria.go)

| Método | Complexidade (média) | Descrição |
|--------|:--------------------:|-----------|
| `Insert(val)` | **O(log n)** | Insere mantendo `esquerda < raiz < direita` |
| `Contains(val)` | **O(log n)** | Busca binária recursiva |
| `Remove(val)` | **O(log n)** | Trata os 3 casos: folha, 1 filho, 2 filhos |
| `InOrder()` | **O(n)** | Retorna valores em ordem crescente |
| `CountLeaves()` | **O(n)** | Conta nós sem filhos |

> **Remoção com 2 filhos:** substitui pelo **sucessor in-order** (menor valor da subárvore direita).

---

## 🚀 Como Executar

Você precisa ter o [Go instalado](https://golang.org/dl/) (versão **1.18+**).

```bash
# 1. Clone o repositório
git clone https://github.com/andrelinodaniel/Estrutura-de-Dados-GO.git
cd Estrutura-de-Dados-GO

# 2. Execute qualquer exemplo diretamente
go run 01-introducao/hello_world.go
go run 07-ordenacao/bubble_sort.go

# 3. Rode as estruturas de dados
go run 08-estruturas-de-dados/pilha/pilha.go
go run 08-estruturas-de-dados/pilha/pilha_parenteses.go
go run 08-estruturas-de-dados/fila/fila_impressao.go
go run 08-estruturas-de-dados/lista-encadeada/lista_duplamente_encadeada.go
go run 08-estruturas-de-dados/arvore-binaria/arvore_binaria.go
```

---

## 💡 Conceitos e Técnicas Utilizadas

| Conceito | Onde é Aplicado |
|----------|----------------|
| **Ponteiros** `*T` e `&var` | Nodes de todas as estruturas encadeadas |
| **Structs com métodos** (receivers) | Pilha, Fila, BST, Lista |
| **Recursividade** | Árvore BST (insert, remove, traverse) e exercícios |
| **Tratamento de erros** com `error` | `Dequeue()` em fila vazia |
| **Maps** em Go | Verificador de parênteses balanceados |
| **Interfaces implícitas** | Modularidade entre pacotes |

---

## 📁 Estrutura Completa

```
Estrutura-de-Dados-GO/
├── 01-introducao/
│   ├── hello_world.go
│   ├── exercicio_1.go
│   └── exercicio_2.go
├── 02-arrays-e-matrizes/
│   ├── arrays/
│   └── matrizes/
├── 03-slices/
├── 04-ponteiros/
├── 05-structs/
│   ├── exercicios-basicos/
│   └── lista-exercicios/       ← exercicio-1 ao exercicio-18
├── 06-recursividade/
│   └── lista-exercicios/
├── 07-ordenacao/
│   ├── bubble_sort.go
│   └── bubble_sort_pratica.go
├── 08-estruturas-de-dados/     ← ⭐ núcleo do repositório
│   ├── README.md
│   ├── pilha/
│   ├── fila/
│   ├── lista-encadeada/
│   └── arvore-binaria/
└── 09-prova/
```

---

<div align="center">

**Andrelino Daniel Filho** · Estudos de Estrutura de Dados em Go · 2026

[![GitHub](https://img.shields.io/badge/GitHub-andrelinodaniel-181717?style=flat-square&logo=github)](https://github.com/andrelinodaniel)

</div>
