package main

import "strings"

var numberPad = [][]string{
	{},                   // 0
	{},                   // 1
	{"a", "b", "c"},      // 2
	{"d", "e", "f"},      // 3
	{"g", "h", "i"},      // 4
	{"j", "k", "l"},      // 5
	{"m", "n", "o"},      // 6
	{"p", "q", "r", "s"}, // 7
	{"t", "u", "v"},      // 8
	{"w", "x", "y", "z"}, // 9
}

func allCases(d string, res *[]string, ch [][]int, count int) {
	var builder strings.Builder

	for i := 0; i < len(d); i++ {
		digitNum := int(d[i] - '0')
		builder.WriteString(numberPad[digitNum][ch[i][0]])
	}
	(*res)[count] = builder.String()

	num := len(ch) - 1
	ok := false
	for num >= 0 {
		if ch[num][0] == ch[num][1]-1 {
			ch[num][0] = 0
			num--
		} else {
			ch[num][0]++
			ok = true
			break
		}
	}
	if ok {
		allCases(d, res, ch, count+1)
	}

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	count := 1

	check := make([][]int, len(digits))

	for i := 0; i < len(digits); i++ {
		check[i] = make([]int, 2)
		digitNum := int(digits[i] - '0')
		check[i][1] = len(numberPad[digitNum])
		count = count * check[i][1]
	}

	res := make([]string, count)

	allCases(digits, &res, check, 0)

	return res
}
