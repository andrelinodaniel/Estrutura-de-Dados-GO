package main

import (
	"errors"
	"fmt"
)

// QueueNode representa um trabalho de impressão na fila.
type QueueNode struct {
	Value int
	Next  *QueueNode
}

// Queue é uma fila FIFO clássica: front = próximo a sair, rear = último que entrou.
// totalProcessed é um contador histórico que nunca diminui.
type Queue struct {
	front          *QueueNode
	rear           *QueueNode
	size           int
	totalProcessed int
}

// Enqueue adiciona um trabalho ao final da fila.
func (q *Queue) Enqueue(value int) {
	newNode := &QueueNode{Value: value}
	if q.rear == nil {
		// fila vazia: o novo nó é front e rear ao mesmo tempo
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
	q.size++
}

// Dequeue remove o trabalho mais antigo (front) e retorna erro se a fila estiver vazia.
func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, errors.New("fila vazia: não há trabalhos para imprimir")
	}
	value := q.front.Value
	q.front = q.front.Next
	if q.front == nil {
		// a fila ficou vazia, então rear também precisa ser resetado
		q.rear = nil
	}
	q.size--
	q.totalProcessed++ // conta como "processado", esse número nunca é decrementado
	return value, nil
}

func (q *Queue) Size() int {
	return q.size
}

// TotalProcessed retorna quantos trabalhos já foram impressos desde o início,
// independente de quantos ainda estão na fila agora.
func (q *Queue) TotalProcessed() int {
	return q.totalProcessed
}

// Clear esvazia a fila atual, mas preserva o histórico (totalProcessed).
func (q *Queue) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()

	fmt.Println("Size:", queue.Size())
	fmt.Println("TotalProcessed:", queue.TotalProcessed())

	queue.Clear()
	fmt.Println("Após Clear -> Size:", queue.Size())
	fmt.Println("Após Clear -> TotalProcessed:", queue.TotalProcessed())

	_, err := queue.Dequeue()
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
