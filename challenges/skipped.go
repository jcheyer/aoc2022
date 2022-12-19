package challenges

type Skipped struct{}

func (d *Skipped) Load(file string) error {
	return nil
}

func (d *Skipped) Part1() string {
	return "skipped"
}

func (d *Skipped) Part2() string {
	return "skipped"
}
