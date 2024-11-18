package main

type Kind byte

const (
	Sell Kind = iota
	Buy
)

func (s Kind) String() string {
	if s == Buy {
		return "buy"
	}

	return "sell"
}
