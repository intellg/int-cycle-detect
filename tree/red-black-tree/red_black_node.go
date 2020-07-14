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

func (node *Node) getParentGrandparent() (parent, grandparent *Node, ok bool) {
	parent = node.Parent
	if parent == nil {
		return
	}
	grandparent = parent.Parent
	if grandparent == nil {
		return
	}
	ok = true
	return
}
