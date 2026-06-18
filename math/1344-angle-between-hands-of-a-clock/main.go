package main

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func angleClock(hour int, minutes int) float64 {
	var h, m, res float64

	m = float64(minutes) * 6
	h = float64(hour)*30 + float64(minutes)*0.5

	res = max(h, m) - min(h, m)
	if res > 180 {
		return 360 - res
	}

	return res
}
