package types

import (
	"time"
)

type Block struct {
	ID              string    `json:"id"`
	Height          int64     `json:"height"`
	Time            time.Time `json:"time"`
	DataHash        string    `json:"data_hash"`
	ValidatorsHash  string    `json:"validators_hash"`
	EvidenceHash    string    `json:"evidence_hash"`
	ProposerAddress string    `json:"proposer_address"`
}
