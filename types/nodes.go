package types

type Speed struct {
	Upload   int
	Download int
}

type Coin struct {
	Amount int
	Denom  string
}

type Coins []Coin

type Node struct {
	ID            int
	Address       string
	Deposit       Coin
	Type          string
	PricesPerGB   Coins
	Version       string
	Moniker       string
	Encryption    string
	InternetSpeed Speed
	AddedAt       string
	UpdatedAt     string
}
