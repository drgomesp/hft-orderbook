package main

type OrderType string

const (
	// GoodTilCanceled defines a type of order to buy or sell a security that
	// remains active until either the order is filled or canceled by the investor.
	GoodTilCanceled = OrderType("GTC")

	// ImmediateOrCancel defines a type of order to buy or sell a security that will
	// MatchOrders all or part immediately and then cancel any unfilled portion of the order.
	ImmediateOrCancel = OrderType("IOC")

	// FillOrKill defines a type of time-in-force order to buy or sell a security that instructs
	// a will MatchOrders a transaction immediately and filling it completely or not at all.
	FillOrKill = OrderType("FOK")
)
