package proccesor

import (
	"testing"
	"time"

	"github.com/arnaz06/deposit/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestThresholdResolver(t *testing.T) {
	now := time.Now()
	tests := []struct {
		testname       string
		history        []*pb.DepositHistory
		currTime       time.Time
		inputtedAmount int64
		expectedResult bool
	}{
		{
			testname:       "first deposit with threshold exceeded",
			currTime:       now,
			inputtedAmount: 11000,
			expectedResult: true,
		},
		{
			testname:       "second deposit with threshold exceeded",
			currTime:       time.Now(),
			inputtedAmount: 6000,
			history: []*pb.DepositHistory{
				{
					Balance:   6000,
					CreatedAt: timestamppb.New(now.Add(-1 * time.Minute)),
				},
			},
			expectedResult: true,
		},
		{
			testname:       "third deposit with threshold exceeded",
			currTime:       time.Now(),
			inputtedAmount: 2000,
			history: []*pb.DepositHistory{
				{
					Balance:   3000,
					CreatedAt: timestamppb.New(now.Add(-30 * time.Second)),
				},
				{
					Balance:   6000,
					CreatedAt: timestamppb.New(now.Add(-30 * time.Second)),
				},
			},
			expectedResult: true,
		},
		{
			testname:       "third deposit with threshold not exceeded",
			currTime:       time.Now(),
			inputtedAmount: 2000,
			history: []*pb.DepositHistory{
				{
					Balance:   3000,
					CreatedAt: timestamppb.New(now.Add(-1 * time.Minute)),
				},
				{
					Balance:   6000,
					CreatedAt: timestamppb.New(now.Add(-2 * time.Minute)),
				},
			},
			expectedResult: false,
		},
		{
			testname:       "third deposit with threshold not exceeded because balance not enough",
			currTime:       time.Now(),
			inputtedAmount: 2000,
			history: []*pb.DepositHistory{
				{
					Balance:   3000,
					CreatedAt: timestamppb.New(now.Add(-30 * time.Second)),
				},
				{
					Balance:   3000,
					CreatedAt: timestamppb.New(now.Add(-30 * time.Second)),
				},
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			res := thresholdResolver(test.history, test.inputtedAmount, test.currTime)
			require.Equal(t, test.expectedResult, res)
		})
	}
}
