package leetcode

// 盛最多水的容器

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，
// 垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 说明：你不能倾斜容器。

//

// 示例 1：

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：

// 输入：height = [1,1]
// 输出：1
// 示例 3：

// 输入：height = [4,3,2,1,4]
// 输出：16
// 示例 4：

// 输入：height = [1,2,1]
// 输出：2

// 计算凹槽， 积水量由最低的决定

type SortMax struct {
	nums  []int       // 从大到小的数组
	index map[int]int // 新索引对应的原数组的索引位置
}

// 排序， 得到一个新的从大到小的数组和对应的索引值, 值必须大于0， 准确来说不能有-1
func Sort(nums []int, asc bool) *SortMax {
	length := len(nums)
	sm := &SortMax{
		nums:  make([]int, 0, length),
		index: make(map[int]int, length),
	}

	for j := 0; j < length; j++ {
		thisIndex := j
		for i := 0; i < length; i++ {
			if nums[i] == -1 {
				continue
			}
			if asc {
				if nums[i] < nums[thisIndex] {
					thisIndex = i
				}
			} else {
				if nums[i] > nums[thisIndex] {
					thisIndex = i
				}
			}
		}
		sm.nums = append(sm.nums, nums[thisIndex])
		sm.index[j] = thisIndex
		nums[thisIndex] = -1
	}

	return sm
}

func (sm *SortMax) GetNums() []int {
	return sm.nums
}

func (sm *SortMax) GetIndexs() map[int]int {
	return sm.index
}

func maxArea(height []int) int {
	// 官方给出的是双指针往中间移动
	// minValue * (l2 - l1) > minValue * (l2 - l1 - 1)  2边的值大于0
	// 也就是小的值往中间移动， 大的值保持不变
	// 如果2个值是一样的，就随意移动一根
	max := 0
	left := 0
	right := len(height) - 1
	// 最大值-1
	for left == right {
		// 2值最小值
		min := height[left]
		if height[left] > height[right] {
			min = height[right]
			right--
		} else {
			left++
		}
		if max < min*(right-left+1) {
			max = min * (right - left + 1)
		}
	}
	return max
}

// func maxArea(height []int) int {
// 	// 数组的长度小于等于1， 无法积水, 此方法会超时
// 	if len(height) <= 1 {
// 		return 0
// 	}
// 	max := 0
// 	for i := range height {
// 		for j := i; j < len(height); j++ {
// 			// 2值最小值
// 			min := height[i]
// 			if height[i] > height[j] {
// 				min = height[j]
// 			}
// 			if max < min*(j-i) {
// 				max = min * (j - i)
// 			}
// 		}
// 	}
// 	return max
// }
