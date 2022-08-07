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
		if len(dep.DepositHistory) == 0 {
			if dep.Balance >= 10000 {
				dep.AboveThreshold = true
			}
		} else {
			diff := dep.DepositHistory[len(dep.DepositHistory)-1].CreatedAt.AsTime().Unix() - time.Now().Unix()
			if diff < 120 && dep.Balance > 10000 {
				dep.AboveThreshold = true
			}
		}
	}

	dep.DepositHistory = append(dep.DepositHistory, &pb.DepositHistory{
		Balance:   dep.Balance,
		CreatedAt: timestamppb.New(time.Now().Add(time.Hour * 1)),
	})
	ctx.SetValue(dep)
	fmt.Printf("success consume data %+v\n", dep)
}
