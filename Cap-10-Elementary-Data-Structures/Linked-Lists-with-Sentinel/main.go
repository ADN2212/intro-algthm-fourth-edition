package main

import "fmt"


//Ver pagina 259 para ver como se modelan los nodos de una lkl doble
type Node struct {
	key int
	prev *Node
	next *Node
}

//Ver pagina 262
//En este caso en vez de una head la lista tiene un sentinela
//el cual sirve para hacer que la lista sea circular, 
//lo cual simplifica de forma considerable todos los algoritmos  
type LinkedList struct {
	sentinel *Node
}

//O(n)
func (lkl *LinkedList) show() {
	curr := lkl.sentinel
	for true {
		fmt.Println(curr.key)
		curr = curr.next
		//Para evitar hacer un bucle infinito
		if curr == lkl.sentinel {
			break
		}
	}
}

//O(1)
func (lkl *LinkedList) delete(node *Node) {
	//En el algoritmo de la pagina 262 no esta este condicional
	//pero en el libro se define una lkl vacia como una que solo tiene su sentinel
	//por eso evito que sea borrable
	if node != lkl.sentinel {
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		fmt.Println("list sentinel can not be deleted")
	}
}

//Ver pag 263
//Ver como la mera existencia del sentinel hace que no sea necesaria la condicional que esta en el algo original
//By the way this is O(1)
func (lkl *LinkedList) insert(newNode, node *Node) {
	newNode.next = node.next
	newNode.prev = node
	node.next.prev = newNode
	node.next = newNode	
}


func (lkl *LinkedList) search(k int) *Node {
	//Primero se guarda la key que se busca en el sentinel para evitar un bucle infinito
	lkl.sentinel.key = k
	curr := lkl.sentinel.next
	
	for curr.key != k {
		curr = curr.next
	}
	//Si al final de la interacion el curr es el sentinel
	//Esto significa que no habia ningun nodo con la key que se buscaba
	if curr == lkl.sentinel {
		return nil
	} else {
		return curr
	}
}


func main() {
	sentinel := Node{key: -100}
	
	//Para una lkl vacia el sentinel es su propio prev y next 
	sentinel.prev = &sentinel
	sentinel.next = &sentinel

	lkl := LinkedList{sentinel: &sentinel}
	//lkl.show()//solo el sentinel
	
	n1 := &Node{key: 9}
	n2 := &Node{key: 16}
	n3 := &Node{key: 4}
	n4 := &Node{key: 1}
	lkl.insert(n1, lkl.sentinel)
	lkl.insert(n2, n1)
	lkl.insert(n3, n2)
	lkl.insert(n4, n3)
	lkl.show()

	fmt.Println("After delete ---------------")
	lkl.delete(n3)
	lkl.show()

	fmt.Println(lkl.search(16))

}
