package types

type Proposals struct {
	ID              int
	Title           string
	Description     string
	Type            string
	Status          bool
	TotalDeposit    Coin
	VotingStartedAt string
	VotingEndedAt   string
}

type Deposits struct {
	Address   string
	PoposalID int
	Amount    Coin
}
