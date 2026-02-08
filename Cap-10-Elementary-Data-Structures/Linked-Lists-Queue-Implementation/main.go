package main

import (
	"fmt"
)

// Ver pagina 259 para ver como se modelan los nodos de una lkl doble
type Node struct {
	key  int
	prev *Node
	next *Node
}

// Se refiere a una lkl doble
type Queue struct {
	front *Node
	back  *Node
}

// O(n) con n igual a la longitud de la lista
func (q *Queue) show() {
	curr := q.front
	for curr != nil {
		fmt.Println(curr.key)
		curr = curr.next
	}
}

// Agrega un alemento al final de la cola
// O(1)
func (q *Queue) enqueue(k int) {
	//Si q no tiene cola significa que esta vacia
	//por lo cual se crea un nuevo nodo que sera la cola y el frente al mismo tiempo
	if q.back == nil {
		q.back = &Node{key: k}
		//Estas asignaciones circulares estan haciendo que el primer nodo se comporte como un sentinel
		q.front = q.back
		q.front.next = q.back
		q.back.prev = q.front
	} else {
		//Si q tiene cola creamos un nuevo nodo
		//y lo ponemos en su lugar correspondiente
		newBack := Node{key: k}
		q.back.next = &newBack
		newBack.prev = q.back
		q.back = &newBack
	}
}

// Este metod solo muestra el nodo que esta en el frente de la cola
func (q *Queue) peak() *Node {
	return q.front
}

// ojo en una implementacion real estos metodos tendrian que estar en mayusculas
// O(1)
func (q *Queue) dequeue() {
	if q.front != nil {
		oldFront := q.front
		var newFront *Node
		//Cuando la queue tiene un solo elemento,
		//el next node de este sera nil,
		//esta condicional sirve para evitar un nil pointer deference error
		if oldFront.next != nil {
			newFront = oldFront.next

		} else {
			newFront = nil
		}
		//Desconecto el viejo de su nodo siguiente y viseversa
		oldFront.next = nil
		if newFront != nil {
			newFront.prev = nil
		}
		//Reasigno el front de la queue
		q.front = newFront
	} else {
		fmt.Println("the queue was empty")
	}
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}

func main() {
	fmt.Println("This is an queue implemented using doble lkls")
	q1 := Queue{}
	q1.enqueue(1)
	q1.enqueue(2)
	q1.enqueue(3)
	q1.show()
	fmt.Println(q1.peak())
	q1.dequeue()
	fmt.Println("After dequeeing -----")
	q1.show()

	//To' junto:

	q2 := Queue{}
	q2.enqueue(20)
	q2.enqueue(25)
	q2.enqueue(30)
	q2.enqueue(35)

	for !q2.isEmpty() {
		fmt.Println(q2.peak().key)
		q2.dequeue()
		//fmt.Println("After dequeue ...")
	}

}
