package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		var buf strings.Builder
		count := 0
		var now byte
		for j := 0; j < len(res); j++ {
			if j == 0 {
				now = res[j]
				count = 1
			} else {
				if now == res[j] {
					count++
				} else {
					buf.WriteString(strconv.Itoa(count))
					buf.WriteString(string(now))
					now = res[j]
					count = 1
				}
			}
		}
		buf.WriteString(strconv.Itoa(count))
		buf.WriteString(string(now))
		res = buf.String()
	}
	return res
}
