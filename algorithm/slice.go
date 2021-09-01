package algorithm

// 合并2个有序的数组为一个新的有序数组， 值可以重复
func MergeTwoOrderSlice(nums1 []int, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 && l2 == 0 {
		return nums1
	} else if l1 == 0 && l2 != 0 {
		return nums2
	}
	if l2 == 0 && l1 != 0 {
		return nums1
	}
	new := make([]int, 0, len(nums1)+len(nums2))
	return mergeTwoOrderSlice(nums1, nums2, new, 0, 0)
}

func mergeTwoOrderSlice(nums1 []int, nums2 []int, new []int, first, second int) []int {

	if nums1[first] <= nums2[second] {
		new = append(new, nums1[first])
		if len(nums1)-1 == first {
			new = append(new, nums2[second:]...)
			return new
		} else {
			first++
		}
	} else {
		new = append(new, nums2[second])
		if len(nums2)-1 == second {
			new = append(new, nums1[first:]...)
			return new
		} else {
			second++
		}
	}
	return mergeTwoOrderSlice(nums1, nums2, new, first, second)
}
