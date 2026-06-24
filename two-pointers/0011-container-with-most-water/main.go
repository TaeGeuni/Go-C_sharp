package main

func maxArea(height []int) int {
	res := 0

	left, right := 0, len(height)-1

	for left < right {
		if res < (right-left)*min(height[right], height[left]) {
			res = (right - left) * min(height[right], height[left])
		}
		if height[left+1] > height[right-1] {
			left++
		} else {
			right--
		}
	}

	return res
}
