package main

func minSubArrayLen(target int, nums []int) int {
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	if total < target {
		return 0
	}
	res := len(nums)
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] < nums[right] {
			total -= nums[left]
			left++
		} else if nums[left] > nums[right] {
			total -= nums[right]
			right--
		} else if nums[left] == nums[right]&&{

		}

		if total < target {
			break
		}
		res--
	}

	return res
}
