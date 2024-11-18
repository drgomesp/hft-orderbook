package main

type Type string

const (
	// GoodTilCanceled defines a type of order to buy or
	// sell a security that remains active until either
	// the order is filled or canceled by the investor.
	GoodTilCanceled = Type("GOOD_TIL_CANCELED")

	// ImmediateOrCancel defines a type of order to buy or
	// sell a security that will execute all or part immediately
	// and then cancel any unfilled portion of the order.
	ImmediateOrCancel = Type("IMMEDIATE_OR_CANCELED")

	// FillOrKill defines a type of time-in-force order to buy
	// or sell a security that instructs a will execute a transaction
	// immediately and filling it completely or not at all.
	FillOrKill = Type("FILL_OR_KILL")
)
