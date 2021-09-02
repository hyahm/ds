package leetcode

// 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

//

// 示例 1：

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：

// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：

// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func Append(ln *ListNode, value int) {
	end := EndNode(ln)
	x := &ListNode{
		Val: value,
	}
	end.Next = x

}

func EndNode(ln *ListNode) *ListNode {
	for {
		if ln.Next == nil {
			return ln
		} else {
			ln = ln.Next
		}
		// 添加元素并返回当前元素
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var jinwei bool
	ln := &ListNode{}
	var count int
	for {
		sum := l1.Val + l2.Val
		if jinwei {
			sum++
			jinwei = false
		}

		if sum >= 10 {
			jinwei = true
			sum = sum - 10
		}
		if count == 0 {
			ln.Val = sum
		} else {
			Append(ln, sum)
		}
		count++
		if l1.Next != nil || l2.Next != nil {
			if l1.Next != nil {
				l1 = l1.Next
			} else {
				l1.Val = 0
			}
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2.Val = 0
			}
		} else {
			if jinwei {
				Append(ln, 1)
			}
			return ln
		}

	}

}
