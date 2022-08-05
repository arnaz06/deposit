package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/pb"
	"github.com/lovoo/goka"
	"google.golang.org/protobuf/proto"
)

type depositRepo struct {
	gokaView *goka.View
	gokaEmit *goka.Emitter
}

type DepositRepository interface {
	Get(ctx context.Context, walletID int64) (deposit.Deposit, error)
	Deposit(ctx context.Context, input *pb.Deposit) error
}

func NewDepositRepository(gokaView *goka.View, gokaEmit *goka.Emitter) DepositRepository {
	return &depositRepo{gokaView: gokaView, gokaEmit: gokaEmit}
}

func (r *depositRepo) Get(ctx context.Context, walletID int64) (deposit.Deposit, error) {
	dep := deposit.Deposit{}
	res, err := r.gokaView.Get(fmt.Sprintf("%d", walletID))
	if err != nil {
		return dep, err
	}
	depByte, err := json.Marshal(res)
	if err != nil {
		return dep, err
	}

	err = json.Unmarshal(depByte, &dep)
	if err != nil {
		return dep, err
	}

	return dep, nil
}

func (r *depositRepo) Deposit(ctx context.Context, input *pb.Deposit) error {
	data, err := proto.Marshal(input)
	if err != nil {
		return err
	}
	err = r.gokaEmit.EmitSync(fmt.Sprintf("%d", &input.WalletId), data)
	if err != nil {
		return err
	}
	return nil
}
