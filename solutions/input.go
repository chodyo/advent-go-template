package solutions

type Input struct{}

func (*Input) SetSeparator(string) {}

func (*Input) Lines() []string { return []string{} }

func (*Input) Text() string { return "" }

func (*Input) Ints() []int { return []int{} }

func (*Input) Int() int { return 0 }
