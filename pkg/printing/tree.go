package printing

import "github.com/silphid/ppds/tree"

type Node struct {
	name     string
	children []*Node
}

func (n *Node) Data() interface{} {
	return n.name
}

func (n *Node) Children() (c []tree.Node) {
	for _, child := range n.children {
		c = append(c, tree.Node(child))
	}
	return
}
