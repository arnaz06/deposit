package proccesor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/arnaz06/deposit/pb"
	"github.com/labstack/gommon/log"
	"github.com/lovoo/goka"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Proccess(ctx goka.Context, msg interface{}) {
	fmt.Println("Processing message:", msg)
	var dep *pb.Deposit
	var savedData *pb.Deposit
	err := json.Unmarshal([]byte(msg.(string)), &dep)
	if err != nil {
		log.Error(err)
	}

	if val := ctx.Value(); val != nil {
		savedData = val.(*pb.Deposit)
		if dep == nil {
			dep = savedData
		}
		dep.Balance += savedData.Balance
		dep.DepositHistory = savedData.DepositHistory
	}

	if !dep.AboveThreshold {
		dep.AboveThreshold = thresholdResolver(dep.DepositHistory, dep.Balance, time.Now())
	}

	dep.DepositHistory = append(dep.DepositHistory, &pb.DepositHistory{
		Balance:   dep.Balance,
		CreatedAt: timestamppb.New(time.Now()),
	})

	ctx.SetValue(dep)
	fmt.Printf("success consume data %+v\n", dep)
}

func thresholdResolver(input []*pb.DepositHistory, inputtedAmount int64, currentTime time.Time) bool {
	if len(input) == 0 {
		return inputtedAmount > 10000
	}

	thresholdTime := int64(120)
	thresholdBalance := int64(10000)
	total := inputtedAmount
	var diff int64
	for i := len(input) - 1; i >= 0; i-- {
		diff += currentTime.Unix() - input[i].CreatedAt.AsTime().Unix()
		if diff > thresholdTime {
			break
		}
		total += input[i].Balance
		if total > thresholdBalance {
			return true
		}
	}

	return false
}
