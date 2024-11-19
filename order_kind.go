package main

type OrderKind byte

const (
	Sell OrderKind = iota
	Buy
)

func (s OrderKind) String() string {
	if s == Buy {
		return "buy"
	}

	return "sell"
}
