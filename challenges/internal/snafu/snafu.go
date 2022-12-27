package snafu

var snafu2dec map[rune]int = map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}

var dec2snafu []string = []string{"0", "1", "2", "=", "-"}

func Snafu2Dec(s string) int {
	res := 0
	for _, digit := range s {
		res = res*5 + snafu2dec[digit]
	}
	return res
}

func Dec2Snafu(i int) string {
	res := ""
	for i > 0 {
		res = dec2snafu[i%5] + res
		if i%5 > 2 {
			i += 5
		}
		i = i / 5
	}
	return res
}
