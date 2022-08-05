package deposit

type Deposit struct {
	WalletID       int64 `json:"wallet_id"`
	Balance        int64 `json:"balance"`
	AboveThreshold bool  `json:"above_threshold"`
}
