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
	fmt.Println("FOOOOOO")

	//Ver grafo en la pagina 322
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

	var node2 Node = Node{
		key:    2,
		parent: &node1,
		left:   nil,
		right:  nil,
	}

	var node3 Node = Node{
		key:    9,
		parent: &node1,
		left:   nil,
		right:  nil,
	}

	//Aqui termina la rama izquierda
	node1.left = &node2
	node1.right = &node3

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

	//Aqui la derecha
	node5.right = &node7

	fmt.Println("Before insertion:")
	inorderTreeWalk(tree1.root)
	insertNode(&tree1, 13)
	insertNode(&tree1, 1)
	insertNode(&tree1, 25)
	fmt.Println("Before insertion:")
	inorderTreeWalk(tree1.root)	


	emptyTree := Tree{
		root: nil,
	}

	//Al insertar un nodo en un empty tree dicho nodo se hace el root del tree
	fmt.Println("Ex-empty tree")
	insertNode(&emptyTree, 10)
	inorderTreeWalk(emptyTree.root)

}

func inorderTreeWalk(root *Node) {
	if root != nil {
		inorderTreeWalk(root.left)
		fmt.Printf("Node(%v) \n", root.key)
		inorderTreeWalk(root.right)
	}
}

//Ver pagina 321
//Esta funcion inserta un nodo dentro de un BST haciendo que se respete la propiead de que 
//node.left.hey <= node.key y node.key <= node.right.key
func insertNode(tree *Tree, val int) {

	currentNode := tree.root
	var parent *Node = nil

	//Este bucle va decendiendo por el arbol tomando en cuenta si el valor del nodo es menor 
	// o >= para hayar el nodo al que debe ser agregadi como hijo.
	for currentNode != nil {
		parent = currentNode
		if val < currentNode.key {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	newNode := Node{
		key: val,
		parent: nil,
		left: nil,
		right: nil,
	}

	//Esto sucede cuando el tree está vacio
	if parent == nil {
		tree.root = &newNode 
	//En caso de que no el arbol no era vacio y parent sera el puntero a uno de los nodos del grafo
	} else if newNode.key < parent.key {	
		parent.left = &newNode
	} else {
		parent.right = &newNode
	}

}

