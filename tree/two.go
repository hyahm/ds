package tree

// 双叉树

type TwoTreeNode struct {
	Val   int
	Left  *TwoTreeNode
	Right *TwoTreeNode
}

func buildTree(nums []int, index int) *TwoTreeNode {
	// index: 当前元素的数组索引
	if index > len(nums)-1 {
		return nil
	}

	node := &TwoTreeNode{
		Val: nums[index],
	}
	node.Left = buildTree(nums, index*2+1)
	node.Right = buildTree(nums, index*2+2)
	return node
}

func sortedArrayToBST(nums []int) *TwoTreeNode {
	return buildTree(nums, 0)
}
