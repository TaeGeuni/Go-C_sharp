package main

func processStr(s string) string {
	btSli := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(btSli) != 0 {
				btSli = btSli[:len(btSli)-1]
			}
		} else if s[i] == '#' {
			btSli = append(btSli, btSli...)
		} else if s[i] == '%' {
			rvSli := make([]byte, len(btSli))
			for j := len(btSli) - 1; j >= 0; j-- {
				rvSli[len(btSli)-1-j] = btSli[j]
			}
			btSli = rvSli
		} else {
			btSli = append(btSli, s[i])
		}
	}

	return string(btSli)
}
