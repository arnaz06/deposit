package service

import (
	"context"

	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/pb"
	"github.com/arnaz06/deposit/repository"
)

type depositService struct {
	repo repository.DepositRepository
}

type DepositService interface {
	Get(ctx context.Context, walletID int64) (deposit.Deposit, error)
	Deposit(ctx context.Context, input deposit.Deposit) error
}

func NewDepositService(repo repository.DepositRepository) DepositService {
	return &depositService{repo: repo}
}

func (s *depositService) Get(ctx context.Context, walletID int64) (deposit.Deposit, error) {
	return s.repo.Get(ctx, walletID)
}

func (s *depositService) Deposit(ctx context.Context, input deposit.Deposit) error {
	depoPb := &pb.Deposit{
		WalletId: input.WalletID,
		Balance:  input.Balance,
	}
	return s.repo.Deposit(ctx, depoPb)
}
