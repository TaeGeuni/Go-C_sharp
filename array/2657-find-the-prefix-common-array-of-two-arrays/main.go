package main

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)

	freq := make([]bool, n+1)
	num := 0

	for i := 0; i < n; i++ {
		if !freq[A[i]] {
			freq[A[i]] = true
		} else {
			num++
		}
		if !freq[B[i]] {
			freq[B[i]] = true
		} else {
			num++
		}
		res[i] = num
	}

	return res
}
