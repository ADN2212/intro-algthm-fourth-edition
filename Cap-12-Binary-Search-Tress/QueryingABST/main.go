package main

import "fmt"

//Esta es la forma en que se definen los nodos de un BST, ver pag 315
type Node struct {
	key int
	parent *Node
	left *Node
	right *Node
}


func main() {
	//Ver el arbol a en la pag 313
	treeHead1 := Node{
		key: 6,
		parent: nil,
		left: nil,
		right: nil,
	}

	
	node2 := Node{
		key: 5,
		parent: &treeHead1,
		left: nil,
		right: nil,
	}

	node3 := Node{
		key: 2,
		parent: &node2,
		left: nil,
		right: nil,
	}

	node2.left = &node3

	node4 := Node{
		key: 5,
		parent: &node2,
		left: nil,
		right: nil,
	}

	//Aqui termina el left subtree del root node
	node2.right = &node4
	treeHead1.left = &node2

	node5 := Node{
		key: 7,
		parent: &treeHead1,
		left: nil,
		right: nil,
	}

	node6 := Node{
		key: 8,
		parent: &node5,
		left: nil,
		right: nil,
	}

	node5.right = &node6
	treeHead1.right = &node5

	inorderTreeWalk(&treeHead1)
	fmt.Println(treeSearch(&treeHead1, 8))
	fmt.Println(iteritiveTreeSearch(&treeHead1, 5))

	


}


//Este metodo muestra los valores de un BST en forma ordenada, asumiendo que el arbolcumple con la propiedad
//de que node.key >= node.left.hey y node.rigt.key >= node.key 
func inorderTreeWalk(root *Node) {
	if root != nil {
		inorderTreeWalk(root.left)
		fmt.Printf("Node(%v) \n", root.key)
		inorderTreeWalk(root.right)
	}
}

//Tambien asumiendo la propiedad antes dicha, esta funcion se encarga de hallar el PRIMER nodo que tenga la key dada
func treeSearch(root *Node, key int) *Node {
	if root == nil || key == root.key {
		return root
	}

	if key < root.key {
		return treeSearch(root.left, key)
	}

	return treeSearch(root.right, key)
}

//Esta es la vversion iteratica de la funcion anterior:
func iteritiveTreeSearch(root *Node, key int) *Node {
	
	//Esto lo hago por que lo que entra como argumento es un ptr y no quiero reasignarlo
	currentNode := root
	
	for currentNode != nil && key != currentNode.key {
		if key < currentNode.key {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return currentNode

}

