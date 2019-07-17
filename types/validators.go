package types

type Delegation struct {
	Address string
	Amount  string
}

type Commission struct {
	CommissionMaxRate       string
	COmmissionMaxChangeRate string
	CommissionRate          string
}

type Profile struct {
	Website     string
	Description string
	Identity    string
}

type Validator struct {
	OwnerAddr         string
	OperatorAddr      string
	ConsusensusPubKey string
	Moniker           string
	VotingPower       int
	BondedToken       string
	Delegations       []Delegation
	BlockID           string
	CommisionParams   Commission
	Profile           Profile
	AddedAt           string
}
