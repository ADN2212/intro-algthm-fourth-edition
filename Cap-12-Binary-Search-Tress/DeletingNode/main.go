package main

import "fmt"

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

//Esto es para poder representar un empty tree
//Un empty tree es un tree que no tiene root node
type Tree struct {
	root *Node
}

func main() {

	fmt.Println("Node deletion")
	//Ver grafo en la pagina 322 (con algunas modificaciones para abarcar los 4 casos)
	var head Node = Node{
		key:    12,
		parent: nil,
		left:   nil,
		right:  nil,
	}

	var tree1 Tree = Tree{
		root: &head,
	}

	var node1 Node = Node{
		key:    5,
		parent: &head,
		left:   nil,
		right:  nil,
	}

	head.left = &node1

	var node3 Node = Node{
		key:    9,
		parent: &node1,
		left:   nil,
		right:  nil,
	}

	
	var node2 Node = Node{
		key:    8,
		parent: &node3,
		left:   nil,
		right:  nil,
	}

	//Aqui termina la rama izquierda
	node1.right = &node3
	node3.left = &node2

	node4 := Node{
		key: 18,
		parent: &head,
		left: nil,
		right: nil,
	}
	
	head.right = &node4
	
	node5 := Node{
		key: 15,
		parent: &node4,
		left: nil,
		right: nil,
	}

	node6 := Node{
		key: 19,
		parent: &node4,
		left: nil,
		right: nil,
	}

	node4.left = &node5
	node4.right = &node6

	node7 := Node{
		key: 17,
		parent: &node5,
		left: nil,
		right: nil,
	}

	node8 := Node{
		key: 13,
		parent: &node5,
		left: nil,
		right: nil,
	}

	//Aqui la derecha
	node5.right = &node7
	node5.left = &node8
	//Ojo: no ejecutes todos los casos juntos porque este puede dar resultados inesperados:
	//Caso (a):
	// fmt.Println("original tree walk: ")
	// inorderTreeWalk(tree1.root)
	// deleteNode(&tree1, &node1)//Node(5)
	// fmt.Println("After deletion: ")
	// inorderTreeWalk(tree1.root)

	//Caso (b):
	// fmt.Println("original tree walk: ")
	// inorderTreeWalk(tree1.root)
	// deleteNode(&tree1, &node3)//Node(9)
	// fmt.Println("After deletion: ")
	// inorderTreeWalk(tree1.root)
	
	//Caso (c)
	// fmt.Println("original tree walk: ")
	// inorderTreeWalk(tree1.root)
	// deleteNode(&tree1, &node5)//Node(15)
	// fmt.Println("After deletion: ")
	// inorderTreeWalk(tree1.root)
	
	// Caso (d)
	fmt.Println("original tree walk: ")
	inorderTreeWalk(tree1.root)
	deleteNode(&tree1, &head)//Node(12) => this one is also the head
	fmt.Println("After deletion: ")
	inorderTreeWalk(tree1.root)
	
}

//Ver pag 323 y 325
//Esta funcion abarca todos los casos necesarios para eliminar un nodo de un BST
func deleteNode(tree *Tree, toDeleteNode *Node) {

	if toDeleteNode.left == nil {
		transplant(tree, toDeleteNode, toDeleteNode.right)//caso (a) pagina 323
	} else if toDeleteNode.right == nil {
		transplant(tree, toDeleteNode, toDeleteNode.left)//caso (b)
	} else {
		//El sucesor del nodo a borrar es el minimo del right sub-tree, PIENSALO!
		toDeleteNodeSuccesor := findTreeMinimun(toDeleteNode.right)
		//Todo este if es la primera parte del caso (d)
		if toDeleteNodeSuccesor != toDeleteNode.right {
			transplant(tree, toDeleteNodeSuccesor, toDeleteNodeSuccesor.right)
			toDeleteNodeSuccesor.right = toDeleteNode.right
			toDeleteNodeSuccesor.right.parent = toDeleteNodeSuccesor
		}
		//Y esta parte es el caso (c) y la segunda parte del caso (d)
		//Toda esta parte se esta encargando de poner el sucesor en la posicion del nodo actual
		transplant(tree, toDeleteNode, toDeleteNodeSuccesor)
		toDeleteNodeSuccesor.left = toDeleteNode.left
		toDeleteNodeSuccesor.left.parent = toDeleteNodeSuccesor
	}
}

//Esta funcion pone el primer nodo en la posicion del segundo
//ver pagina 324 
//Notece como este metodo no cambia los punteros de currentNode, pero si los de su padre
func transplant(tree *Tree, currentNode *Node, toChangeNode *Node) {

	if currentNode.parent == nil {
		//Si el nodo actaul no tiene padre entonces es el root del tree
		tree.root = toChangeNode
	} else if currentNode == currentNode.parent.left {
		//Si el current node es el left child de su padre lo cambiamos por toChange
		currentNode.parent.left = toChangeNode
	} else {
		//lo mismo pero con el right
		currentNode.parent.right = toChangeNode
	}

	//Finalmente asginamos el parent de el nodo actaul al nodo a cambiar
	//Esto es para evitar un nil pointer exception ???
	if toChangeNode != nil {
		toChangeNode.parent = currentNode.parent
	}

}

func findTreeMinimun(root *Node) *Node {
	currentNodePrt := root
	
	for currentNodePrt.left != nil {
		currentNodePrt = currentNodePrt.left
	}

	return currentNodePrt
}

func inorderTreeWalk(root *Node) {
	if root != nil {
		inorderTreeWalk(root.left)
		fmt.Printf("Node(%v) \n", root.key)
		inorderTreeWalk(root.right)
	}
}
