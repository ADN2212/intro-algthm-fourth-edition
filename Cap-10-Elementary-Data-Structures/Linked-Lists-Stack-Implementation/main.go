package main

import "fmt"

//Sengun mi vision
//Cada nodo de un stack solo debe saber quien esta antes de el
type Node struct {
	key  int
	prev *Node
}

type Stack struct {
	top *Node
}

//Notece como esta operacion tambien es O(1)
func (stk *Stack) push(k int) {
	if stk.top == nil {
		stk.top = &Node{key: k}
	} else {
		newTop := Node{key: k}
		newTop.prev = stk.top
		stk.top = &newTop
	}
}

//Este metodo muestra un stack desde arriba hacia abajo
func (stk *Stack) show() {
	curr := stk.top
	for curr != nil {
		fmt.Println(curr.key)
		curr = curr.prev
	}
}

func (stk *Stack) pop() *Node {
	if stk.top != nil {
		oldTop := stk.top
		stk.top = oldTop.prev
		return oldTop
	} else {
		return nil//Si el metodo retorna nil significa que la stack estava vacia
	}
}

func (stk *Stack) isEmpty() bool {
	return  stk.top == nil
}

func main() {
	fmt.Println("Let's do some shit wiht an stack")

	stk1 := Stack{} //An empty stack
	stk1.push(10)
	stk1.push(9)
	stk1.push(8)
	stk1.show()

	//Se puede usar el metodo pop para recorrer el stack de arriba pa' bajo
	curr := stk1.pop()
	for curr != nil {
		fmt.Println("poping ...")
		fmt.Println(curr.key)
		curr = stk1.pop()
	}

	fmt.Println(stk1.isEmpty())


}
