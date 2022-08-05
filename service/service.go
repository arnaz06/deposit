package service

import (
	"context"

	"github.com/arnaz06/deposit"
)

type DepositService interface {
	Get(ctx context.Context, walletID int64) (deposit.Deposit, error)
	Deposit(ctx context.Context, input deposit.Deposit) (deposit.Deposit, error)
}
