package algo

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n == nil {
		return
	} else if val <= n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Walk() string {
	return ""
}

func (n *Node) Delete(val int) {

}

func (n *Node) InsertAll(array []int) {
	for i := 0; i < len(array); i++ {
		n.Insert(array[i])
	}
}
