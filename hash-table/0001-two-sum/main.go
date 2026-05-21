package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if j, ok := numMap[complement]; ok {
			res[0], res[1] = i, j
			break
		}
		numMap[nums[i]] = i
	}

	return res
}
