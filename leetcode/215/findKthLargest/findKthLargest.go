package findKthLargest

func FindKthLargest(nums []int, k int) int {
	i := 0
	j := len(nums) - 1
	a := nums[0]
	for i < j {
		if nums[j] >= a && j > i {
			j--
		}
		nums[i] = nums[j]
		if nums[i] <= a && j > i {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = a
	// 231589
	if len(nums)-i == k {
		return nums[i]
	} else if len(nums)-i > k {
		return FindKthLargest(nums[i+1:], k)
	} else {
		return FindKthLargest(nums[:i], k-(len(nums)-i))
	}
}
