package redBlackTree

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @Author : 庄广
 * @File : redBlackTree_test
 * @Time : 2021/5/3 20:38
 * @Software: GoLand
 */
//
func TestRedBlackTree(t *testing.T) {
	// [1,5,9,1,5,9]
	// 2
	// 3
	// [1,2,5,6,7,2,4]
	// 4
	// 0
	// [4,1,6,3]
	// 100
	// 1

	fmt.Println(containsNearbyAlmostDuplicate([]int{1,2,3,4,5,6,8,7,11,9,10,13,14,15,16,26,17,18,19,21,22,23,25,24,38,27,28,29,30,31,32,33,34,35,36,37}, 10, 0))
}

type Val int

func (item Val) LessThan(other interface{}) bool {
	o := int(other.(Val))
	if int(item) <= o {
		return true
	}
	return false
}

//https://leetcode-cn.com/problems/contains-duplicate-iii/
//用此题来检验红黑树
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	tree := NewRedBlackTree()
	nodes := make([]*Node, k+1)
	z := 0
	for i:=0; i<len(nums); i++ {
		node := new(Node)
		node.Val = Val(nums[i])
		tree.Insert(node)
		if nodes[z] != nil {
			tree.Delete(nodes[z])
		}
		pred := tree.Predecessor(node)
		if pred != tree.Nil {

			if t >= int(math.Abs(float64(int(pred.Val.(Val))-int(node.Val.(Val))))) {
				return true
			}
		}

		succ := tree.Successor(node)
		if succ != tree.Nil {
			if t >= int(math.Abs(float64(int(succ.Val.(Val))-int(node.Val.(Val))))) {
				return true
			}
		}

		nodes[z] = node
		z++
		z = z%(k+1)
	}
	return false
}
