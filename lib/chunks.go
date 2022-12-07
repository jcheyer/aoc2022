package lib

func Chunks(data []string) [][]string {

	ret := make([][]string, 0)
	chunk := make([]string, 0)
	for _, l := range data {
		if l == "" {
			ret = append(ret, chunk)
			chunk = make([]string, 0)
			continue
		}
		chunk = append(chunk, l)
	}
	ret = append(ret, chunk)

	return ret
}
