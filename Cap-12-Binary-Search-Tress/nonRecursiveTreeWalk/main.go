package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

func main() {

	fmt.Println("Non-recurisve ttree walk ...")

	//Ver el arbol a en la pag 313
	treeHead1 := Node{
		key: 10,
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
		key: 9,
		parent: &node2,
		left: nil,
		right: nil,
	}

	//Aqui termina el left subtree del root node
	node2.right = &node4
	treeHead1.left = &node2

	node5 := Node{
		key: 15,
		parent: &treeHead1,
		left: nil,
		right: nil,
	}

	node6 := Node{
		key: 20,
		parent: &node5,
		left: nil,
		right: nil,
	}

	node5.right = &node6
	treeHead1.right = &node5

	t1 := Tree{
		root: &treeHead1,
	}

	fmt.Println("Recurisve ttree walk ...")
	inorderTreeWalk(t1.root)
	fmt.Println("Non-recurisve ttree walk ...")
	nonRecursiveTreeWalk(&t1)
	
}

func nonRecursiveTreeWalk(tree *Tree) {

	currNode := tree.root
	stack := []*Node{currNode}//Go no tiene buidl in slices

	for currNode != nil {
		if currNode.left == nil && currNode.right == nil {
		 	last, newStack := pop(stack)
			//Debe haber una manera de que este condicional asqueroso no este aqui XD
			if last == nil {
				break
			}

			stack = newStack//this is just reasining a pointer
			showNode(currNode)
			disconect(currNode)
			currNode, newStack = pop(stack)
			stack = newStack
		} else if currNode.left != nil {
			stack = append(stack, currNode.left)
			currNode = currNode.left
		} else {
			//fmt.Println("3")
			showNode(currNode)
			disconect(currNode)
			stack = append(stack, currNode.right)
			currNode = currNode.right
		}
	}
}


func disconect(node *Node) {
	if node != nil && node.parent != nil {
		if node.parent.left == node {
			node.parent.left = nil
		}
		if node.parent.right == node {
			node.parent.right = nil
		}
		node.parent = nil
	}
}

func showNode(node *Node) {
	if node == nil {
		fmt.Println("Nil node")
	} else {
		//fmt.Println("------------------")
		fmt.Printf("Node(%v)\n", node.key)
		//fmt.Println("------------------")
	}
}


func pop(stack []*Node) (*Node, []*Node) {
	if len(stack) == 0 {
		return nil, []*Node{}
	}

	last := stack[len((stack)) - 1]
	newStack := stack[0: len(stack) - 1]

	return last, newStack

}

func inorderTreeWalk(root *Node) {
	if root != nil {
		inorderTreeWalk(root.left)
		fmt.Printf("Node(%v) \n", root.key)
		inorderTreeWalk(root.right)
	}
}





