package main

func numberOfSubstrings(s string) int {
	res := 0
	var check [3]int
	left := 0
	for right := 0; right < len(s); right++ {
		if s[right] == 'a' {
			check[0]++
		} else if s[right] == 'b' {
			check[1]++
		} else if s[right] == 'c' {
			check[2]++
		}

		for check[0] > 0 && check[1] > 0 && check[2] > 0 {
			res += len(s) - right
			if s[left] == 'a' {
				check[0]--
			} else if s[left] == 'b' {
				check[1]--
			} else if s[left] == 'c' {
				check[2]--
			}
			left++
		}

	}

	return res
}
