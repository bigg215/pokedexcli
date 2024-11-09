package main

type ballType int

const (
	PokeBall ballType = iota
	GreatBall
	SafariBall
	MasterBall
)

func (b ballType) String() string {
	return [...]string{
		"poke ball",
		"great ball",
		"safari ball",
		"master ball",
	}[b]
}

func (b ballType) EnumIndex() int {
	return int(b)
}

func (b ballType) Prob() int {
	return [...]int{
		0,
		25,
		75,
		255,
	}[b]
}
