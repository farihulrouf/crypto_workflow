package ledger

type CreateLedgerRequest struct {
	Asset  string  `json:"asset"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
