package types

type Subscription struct {
	ID                 int
	NodeID             int
	Address            string
	PricePerGB         Coin
	TotalDeposit       Coin
	RemainingDeposit   Coin
	RemainingBandwidth Speed
	Status             string
	AddedAt            string
	UpdatedAt          string
}
