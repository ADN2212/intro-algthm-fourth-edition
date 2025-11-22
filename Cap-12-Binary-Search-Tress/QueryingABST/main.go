package main

import (
	"fmt"
)

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

	//Ver grafo en la pag 317
	treeHead2 := Node{
		key: 15,
		parent: nil,//Toda root/head node tine como parent a nil, es decir, no tiene padre.
		left: nil,
		right: nil,
	}

	node7 := Node{
		key: 6,
		parent: &treeHead2,
		left: nil,
		right: nil,
	}

	treeHead2.left = &node7

	node8 := Node{
		key: 3,
		parent: &node7,
		left: nil,
		right: nil,
	}

	node9 := Node{
		key: 7,
		parent: &node7,
		left: nil,
		right: nil,
	}

	node7.left = &node8
	node7.right = &node9

	node10 := Node{
		key: 2,
		parent: &node8,
		left: nil,
		right: nil,
	}

	node11 := Node{
		key: 4,
		parent: &node8,
		left: nil,//Si estos valores no son reasignados entonces esta es una leaf
		right: nil,	
	}

	node8.left = &node10
	node8.right = &node11

	node12 := Node{
		key: 13,
		parent: &node9,
		left: nil,
		right: nil,
	}

	node13 := Node{
		key: 9,
		parent: &node12,
		left: nil,
		right: nil,
	}

	node9.right = &node12
	node12.left = &node13//Aqui termina la rama izquierda


	//Rama derecha:
	node14 := Node{
		key: 18,
		parent: &treeHead2,
		left: nil,
		right: nil,
	}

	node15 := Node{
		key: 17,
		parent: &node14,
		left: nil,
		right: nil,
	}

	node16 := Node{
		key: 20,
		parent: &node14,
		left: nil,
		right: nil,
	}

	node14.left = &node15
	node14.right = &node16
	//Fin de la rama derecha y del grafo.

	treeHead2.right = &node14
	fmt.Println("------------------------------------------------")
	inorderTreeWalk(&treeHead2)
	fmt.Println(treeSearch(&treeHead2, 13))
	fmt.Println(iteritiveTreeSearch(&treeHead2, 4))
	fmt.Println(treeSearch(&treeHead2, 150))
	fmt.Println(findTreeMinimun(&treeHead2))
	fmt.Println(findTreeMinimunRecursive(&treeHead2))
	fmt.Println(findTreeMaximunRecursive(&treeHead2))
	fmt.Println(findTreeMaximun(&treeHead2))

	fmt.Println("------------------------------------------------")
	fmt.Println(findSuccessor(&node12))
	fmt.Println(findSuccessor(&node15))

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

//Si el BST comple con la propiedad antes dicha para hallar el nodo con menor valor solo hay que
//Buscar en profundidad el nodo que este mas a la izquerda
func findTreeMinimun(root *Node) *Node {
	currentNodePrt := root
	
	for currentNodePrt.left != nil {
		currentNodePrt = currentNodePrt.left
	}

	return currentNodePrt
}

func findTreeMinimunRecursive(root *Node) *Node {
	if root.left == nil {
		return root
	}

	return findTreeMinimunRecursive(root.left)
}

//Siguiendo el mismo principio, pera hallar el mayor valor solo hay que buscar en profundidad hacia la derecha
func findTreeMaximunRecursive(root *Node) *Node {
	if root.right == nil {
		return root
	}
	return  findTreeMaximunRecursive((root.right))
}

func findTreeMaximun(root *Node) *Node {
	currentNodePrt := root
	
	for currentNodePrt.right != nil {
		currentNodePrt = currentNodePrt.right
	}

	return currentNodePrt
}

//El sucesor de un nodo el nodo que esta despue de el en un inorder tree walk, ver primera funcion de este file
func findSuccessor(node *Node) *Node {
	
	if node.right != nil {
		//Si el nodo tiene un sub-tree a la derecha, el sucesor es el minimo de ese subtree
		fmt.Println("By minumun")
		return findTreeMinimun(node.right)
	}

	fmt.Println("De other way")
	sucessor := node.parent
	currentNode := node
	
	//La primera condicion es para detener la iteracion cuando se llegue al root node
	//La sengunda se usa para detectar un "peak" en el sub-tree,
	//Lo cual ocurre cuando el nodo actual es el nodo derecho del paret (igual al sucesor en el walk) 
	for sucessor != nil && currentNode == sucessor.right {
		currentNode = sucessor
		sucessor = sucessor.parent//Esta linea hace que se escale en el tree 
	}

	return sucessor


}

