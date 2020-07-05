package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
Root  bool
Data  string
left  *Node
right *Node
}

func (node *Node) insert(value string) error {
newNode := &Node{Data: value}

if node.Root == false {
node.Root = true
node.Data = value
return nil
}

switch strings.Compare(node.Data, value) {
case 1:
if node.left == nil {
node.left = newNode
}
return node.left.insert(value)
case -1:
if node.right == nil {
node.right = newNode
}
return node.right.insert(value)
default:
return errors.New("this value is already in the tree")
}

return nil
}
func (node *Node) search(word string) bool {
current := node

for current != nil {
if current.Data == word {
return true
} else if current.Data > word {
current = current.left
} else if current.Data < word {
current = current.right
}
}
return false

}

func (node *Node) numberLeafNodes() int {
return node.countLeafNodes(node)
}
func (node *Node) countLeafNodes(n *Node) int {
if n == nil {
return 0
}
temp := 0
if n.left == nil && n.right == nil {
temp = 1
}

if n.left != nil {
temp += n.countNodes(n.left)
}
if n.right != nil {
temp += n.countNodes(n.right)

}
return temp
}

func (node *Node) numberNodes() int {
return node.countNodes(node)
}
func (node *Node) countNodes(n *Node) int {
if n == nil {
return 0
}
temp := 1

if n.left != nil {
temp += n.countNodes(n.left)
}
if n.right != nil {
temp += n.countNodes(n.right)

}
return temp
}

func (node *Node) height() int {
return node.findHeight(node) - 1
}

func (node *Node) findHeight(n *Node) int {
if n == nil {
return 0
}

leftHeight := node.findHeight(n.left)
rightHeight := node.findHeight(n.right)

if leftHeight > rightHeight {
return leftHeight + 1
} else {
return rightHeight + 1
}

}

func inOrderTraversal(tree *Node) {

if tree == nil {
return
}

inOrderTraversal(tree.left)
fmt.Println(tree.Data)
inOrderTraversal(tree.right)
}