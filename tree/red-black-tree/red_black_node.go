package red_black_tree

const (
	BLACK = iota
	RED
)

type Node struct {
	Key    int   `json:"key"`
	Color  int   `json:"color"`
	Left   *Node `json:"left"`
	Right  *Node `json:"right"`
	Parent *Node `json:"-"`
}

func (node *Node) Search(key int) *Node {
	n := node
	for n != nil && n.Key != key {
		if key < n.Key {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return n
}

func (node *Node) Minimum() *Node {
	n := node
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (node *Node) Maximum() *Node {
	n := node
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (node *Node) Successor() *Node {
	n := node
	if n.Right != nil {
		return n.Right.Minimum()
	}

	p := n.Parent
	for p != nil && n == p.Right {
		n = p
		p = p.Parent
	}
	return p
}

func (node *Node) Predecessor() *Node {
	n := node
	if n.Left != nil {
		return n.Left.Maximum()
	}

	p := n.Parent
	for p != nil && n == p.Left {
		n = p
		p = p.Parent
	}
	return p
}
