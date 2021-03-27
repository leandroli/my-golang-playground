package findKthLargest_test

import (
	"testing"

	"github.com/leandroli/my-golang-playground/leetcode/215/findKthLargest"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{7,6,5,4,3,2,1}
	if r := findKthLargest.FindKthLargest(nums, 2); r != 6 {
		t.Errorf("FindKthLargest = %d, want 6", r)
	}
}