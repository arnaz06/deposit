package deposit

import "time"

type Deposit struct {
	WalletID       int64            `json:"wallet_id"`
	Balance        int64            `json:"balance"`
	AboveThreshold bool             `json:"above_threshold"`
	DepositHistory []DepositHistory `json:"-"`
}

type DepositHistory struct {
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
