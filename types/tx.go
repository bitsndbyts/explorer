package types

type Tx struct {
	Hash      string `json:"hash" `
	BlockID   string `json:"block_id"`
	Height    int64  `json:"height"`
	Type      string `json:"type"`
	From      string `json:"from"`
	Msg       []byte `json:"msg"`
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Fee       string `json:"fee"`
	Status    bool   `json:"status"`
}
