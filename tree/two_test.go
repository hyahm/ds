package tree

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	ttn := sortedArrayToBST(nums)

	fmt.Println(ttn.Val)
	fmt.Println(ttn.Left.Val)
	fmt.Println(ttn.Right.Val)
	fmt.Println(ttn.Left.Left.Val)
	fmt.Println(ttn.Left.Right.Val)
	fmt.Println(ttn.Right.Left.Val)
	fmt.Println(ttn.Right.Right.Val)
}
