package redBlackTree

/**
 * @Author : 庄广
 * @File : redBlackTree
 * @Time : 2021/4/18 16:05
 * @Software: GoLand
 */

type RedBlackTree struct {
	Root *Node
	Nil  *Node
}

type NodeColor int
const (
	RED NodeColor = iota
	BLACK
)

type Node struct {
	Left *Node
	Right *Node
	Parent *Node
	Val Comparable
	Color NodeColor
}

type Comparable interface {
	LessThan(other interface{}) bool
}

func NewRedBlackTree() *RedBlackTree  {
	nilTree := new(Node)
	nilTree.Color = BLACK
	nilTree.Parent = nilTree
	nilTree.Left = nilTree
	nilTree.Right = nilTree
	tree := &RedBlackTree{
		Root:nilTree,
		Nil:nilTree,
	}
	return tree
}

func (t *RedBlackTree)LeftRotation(x *Node)  {
	//turn y's left subtree into x's right substree
	y := x.Right
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.Parent = x
	}

	//link x's parent to y
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	}else if x == x.Parent.Left {
		x.Parent.Left = y
	}else {
		x.Parent.Right = y
	}

	//put x on y's left
	y.Left = x
	x.Parent = y
}

func (t *RedBlackTree)RightRotation(x *Node)  {
	//turn y's right subtree into x's left substree
	y := x.Left
	x.Left = y.Right
	if y.Right != t.Nil {
		y.Right.Parent = x
	}

	//link x's parent to y
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	}else if x == x.Parent.Right {
		x.Parent.Right = y
	}else {
		x.Parent.Left = y
	}

	//put x on y's right
	y.Right = x
	x.Parent = y
}

func (t *RedBlackTree)Insert(z *Node)  {
	p := t.Nil
	x := t.Root
	for x != t.Nil {
		p = x
		if z.Val.LessThan(x.Val) {
			x = x.Left
		}else {
			x = x.Right
		}
	}
	z.Parent = p
	if p == t.Nil {
		t.Root = z
	}else if z.Val.LessThan(p.Val) {
		p.Left = z
	}else {
		p.Right = z
	}
	z.Left = t.Nil
	z.Right = t.Nil
	z.Color = RED
	t.InsertFixup(z)
}

func (t *RedBlackTree)InsertFixup(z *Node)  {
	for z.Parent.Color == RED {
		p := z.Parent
		grandParent := z.Parent.Parent
		if p == grandParent.Left {
			uncle := grandParent.Right
			if uncle.Color == RED {
				p.Color = BLACK
				uncle.Color = BLACK
				grandParent.Color = RED
				z = grandParent
				continue
			}else if z == p.Right {
				z = p
				t.LeftRotation(z)
			}
			z.Parent.Color = BLACK
			z.Parent.Parent.Color = RED
			t.RightRotation(z.Parent.Parent)

		}else {
			uncle := grandParent.Left
			if uncle.Color == RED {
				p.Color = BLACK
				uncle.Color = BLACK
				grandParent.Color = RED
				z = grandParent
				continue
			}else if z == p.Left {
				z = p
				t.RightRotation(z)
			}
			z.Parent.Color = BLACK
			z.Parent.Parent.Color = RED
			t.LeftRotation(z.Parent.Parent)
		}
	}
	t.Root.Color = BLACK
}


func (t *RedBlackTree)Delete(z *Node)  {
	deleteColorNode := z
	deleteColorNode_original_color := deleteColorNode.Color
	var x *Node
	if z.Left == t.Nil {
		x  = z.Right
		t.Transplant(z, z.Right)
	}else if z.Right == t.Nil {
		x = z.Left
		t.Transplant(z, z.Left)
	}else {
		deleteColorNode = t.TreeMinimum(z.Right)
		deleteColorNode_original_color = deleteColorNode.Color
		x = deleteColorNode.Right
		if deleteColorNode.Parent == z {
			x.Parent = deleteColorNode
		}else {
			t.Transplant(deleteColorNode, deleteColorNode.Right)
			deleteColorNode.Right = z.Right
			deleteColorNode.Right.Parent = deleteColorNode
		}
		t.Transplant(z, deleteColorNode)
		deleteColorNode.Left = z.Left
		deleteColorNode.Left.Parent = deleteColorNode
		deleteColorNode.Color = z.Color
	}

	if deleteColorNode_original_color == BLACK {
		if x.Parent.Right == t.Nil && x.Parent.Left == t.Nil {
			t.DeleteFixup(x)
			return
		}
		t.DeleteFixup(x)
	}
}

func (t *RedBlackTree)DeleteFixup(x *Node)  {
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			right := x.Parent.Right
			if right.Color == RED {
				//case 1 put right.Color = Black
				right.Color = BLACK
				x.Parent.Color = RED
				t.LeftRotation(x.Parent)
				right = x.Parent.Right
			}
			if right.Left.Color == BLACK && right.Right.Color == BLACK {
				//case 2 fixup right tree black color reduce 1
				// count of black in x.Parent reduce 1, x.Parent fixup
				right.Color = RED
				x = x.Parent
				continue
			}else if right.Right.Color == BLACK {
				//case 3 put right.right.color = red
				right.Left.Color = BLACK
				right.Color = RED
				t.RightRotation(right)
				right = x.Parent.Right
			}
			//case 4 if right.right.color = red then left tree black color add 1 fixup finish
			right.Color = x.Parent.Color
			x.Parent.Color = BLACK
			right.Right.Color = BLACK
			t.LeftRotation(x.Parent)
			x = t.Root
		}else {
			//mirror the right,
			left := x.Parent.Left
			if left.Color == RED {
				left.Color = BLACK
				x.Parent.Color = RED
				t.RightRotation(x.Parent)
				left = x.Parent.Left
			}
			if left.Left.Color == BLACK && left.Right.Color == BLACK {
				left.Color = RED
				x = x.Parent
				continue
			}else if left.Left.Color == BLACK {
				left.Right.Color = BLACK
				left.Color = RED
				t.LeftRotation(left)
				left = x.Parent.Left
			}
			left.Color = x.Parent.Color
			x.Parent.Color = BLACK
			left.Left.Color = BLACK
			t.RightRotation(x.Parent)
			x = t.Root
		}
	}
	x.Color = BLACK
}

//replace替换orgin,
func (t *RedBlackTree)Transplant(origin *Node, replace *Node)  {
	if origin.Parent == t.Nil {
		t.Root = replace
	}else if origin == origin.Parent.Left {
		origin.Parent.Left = replace
	}else {
		origin.Parent.Right = replace
	}
	replace.Parent = origin.Parent

}


//指定子树最大值节点
func (t *RedBlackTree)TreeMinimum(n *Node) *Node {
	for n.Left != t.Nil {
		n = n.Left
	}
	return n
}

func (t *RedBlackTree)TreeMaximum(n *Node) *Node {
	for n.Right != t.Nil {
		n = n.Right
	}
	return n
}


//用于按val大小遍历
//指点节点后继
func (t *RedBlackTree)Successor(x *Node) *Node {
	if x.Right != t.Nil {
		return t.TreeMinimum(x.Right)
	}
	succ := x.Parent
	for succ != t.Nil && x == succ.Right {
		x = succ
		succ = succ.Parent
	}
	return succ
}

//指定节点前驱
func (t *RedBlackTree)Predecessor(x *Node) *Node {
	if x.Left != t.Nil {
		return t.TreeMaximum(x.Left)
	}
	pred := x.Parent
	for pred != t.Nil && x == pred.Left {
		x = pred
		pred = pred.Parent
	}
	return pred
}
