package main

func maxNumberOfBalloons(text string) int {
	check := make([]int, 5)

	for i := 0; i < len(text); i++ {
		if text[i] == 'b' {
			check[0]++
		} else if text[i] == 'a' {
			check[1]++
		} else if text[i] == 'l' {
			check[2]++
		} else if text[i] == 'o' {
			check[3]++
		} else if text[i] == 'n' {
			check[4]++
		}
	}
	res := check[0]
	for i := 0; i < 5; i++ {
		if i == 2 || i == 3 {
			if res > check[i]/2 {
				res = check[i] / 2
			}
		} else {
			if res > check[i] {
				res = check[i]
			}
		}
	}
	return res
}
