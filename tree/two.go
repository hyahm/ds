package tree

// 双叉树

type TwoTreeNode struct {
	Val   interface{}
	Left  *TwoTreeNode
	Right *TwoTreeNode
}

func buildTree(nums []interface{}, index int, level int) *TwoTreeNode {
	// index: 当前元素的数组索引
	// css:  层级
	if index > len(nums)-1 {
		return nil
	}

	node := &TwoTreeNode{
		Val: nums[index],
	}
	level++
	node.Left = buildTree(nums, index*2+1, level)
	node.Right = buildTree(nums, index*2+2, level)
	return node
}

func sortedArrayToBST(nums []interface{}) *TwoTreeNode {
	return buildTree(nums, 0, 0)
}
