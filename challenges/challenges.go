package challenges

type Challenge interface {
	Load(input string) error
	Part1() string
	Part2() string
}
