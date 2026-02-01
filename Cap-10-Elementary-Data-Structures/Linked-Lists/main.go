package main

import "fmt"

//Ver pagina 259 para ver como se modelan los nodos de una lkl doble
type Node struct {
	key int
	prev *Node
	next *Node
}

//Se refiere a una lkl doble
type LinkedList struct {
	head *Node
}

//Esto es un metodo no una funcion suelta
func (lkl *LinkedList) show() {
	curr := lkl.head
	for curr != nil {
		fmt.Println(curr.key)
		curr = curr.next
	}
}

//Primer algoritmo, buscar un nodo que contenga una key en una lkl
func (lkl *LinkedList) search(k int) *Node {
	curr := lkl.head
	for curr != nil && k != curr.key {
		curr = curr.next
	}
	//el resultado sera nil si la lkl esta vacia o no se hallo la k
	return curr
}

//Segundo algoritmo, insertar un nodo al principio de una lista:
func (lkl *LinkedList) prepend(newHead *Node) {
	newHead.next = lkl.head
	newHead.prev = nil//por definicion ha de ser asi
	if lkl.head != nil {
		//como la cabeza puede no estar definida
		//Hacemos ese check antes de hacer que la cabeza antigia sea el nodo siguiente de la nueva head
		lkl.head.prev = newHead
	}
	lkl.head = newHead
}

//Tercer algoritmo, insertar un nodo delante de otro
//Ojo, creo que este algoritmo es una generalizacion de prepend
func (lkl *LinkedList) insert(newNode, node *Node) {
	//Siendo node el node que ya estaba en la lkl
	//hacemos su next sea el next del nevo
	//y lo mismo con el prev
	newNode.next = node.next
	newNode.prev = node.prev
	if node.next != nil {
		//Hacemos que le prev del siguiente sea nuestro nuevo nodo
		node.next.prev = newNode
	}
	node.next = newNode
}

//Por ultimo delete
//A pesar de su nombre este algoritmo no elimina el nodo
//Si no que lo quita de la lista al desconectar sus punteros 
func (lkl *LinkedList) delete(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next 
	} else {
		//Si el prev del nodo es nil, entonces es la head
		lkl.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

func main() {
	fmt.Println("Let's do some shit using lkls")

	//(a) de la pagina 259
	h1 := Node{prev: nil, key: 9, next: nil}
	n1 := Node{prev: &h1, key: 16, next: nil}
	h1.next = &n1
	n2 := Node{prev: &n1, key: 4, next: nil}
	n1.next = &n2 
	n3 := Node{prev: &n2, key: 1, next: nil}
	n2.next = &n3
	lkl1 := LinkedList{head: &h1}

	//lkl1.show()	

	//Cambia el valor para testear
	res := lkl1.search(100)
	if res != nil {
		println(res.key)
	} else {
		println("Not found")
	}

	lkl1.show()
	
	println("------------")
	lkl1.prepend(&Node{key: 25})//Notece como Go permite no asignar los valores que son punteros, ya que son nullables
	lkl1.show()

	println("------------")
	lkl1.insert(&Node{key: 8}, &n1)
	lkl1.show()

	println("------------")
	lkl1.delete(&n1)
	lkl1.show()


}



