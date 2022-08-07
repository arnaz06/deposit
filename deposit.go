package deposit

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/arnaz06/deposit/pb"
)

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

type DepositDecoder struct{}

func (dc *DepositDecoder) Encode(value interface{}) ([]byte, error) {
	if v, ok := value.(*pb.Deposit); ok {
		return json.Marshal(v)
	}
	return []byte(""), errors.New("invalid type")
}

func (dc *DepositDecoder) Decode(data []byte) (interface{}, error) {
	var deposit pb.Deposit
	err := json.Unmarshal(data, &deposit)
	if err != nil {
		return nil, err
	}
	return &deposit, nil
}
