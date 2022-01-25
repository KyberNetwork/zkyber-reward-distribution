package cmd

import (
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/reward"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/util/env"
	"github.com/spf13/cobra"
	"math"
)

func newRewardCmd() *cobra.Command {
	var (
		cycle          int
		startTimestamp int
		endTimestamp   int
		totalReward    string
	)

	var command = &cobra.Command{
		Use:   "reward",
		Short: "Calculate reward based on users list files",
		RunE: func(cmd *cobra.Command, args []string) error {
			kyberReward := reward.NewKyberReward(uint64(startTimestamp), uint64(endTimestamp), totalReward)

			return kyberReward.CalculateRewardForUsers(cycle)
		},
	}

	command.Flags().IntVar(&cycle, "cycle", env.ParseNumFromEnv(common.EnvVarCycle, common.DefaultCycle, 0, math.MaxInt), "cycle required")
	command.Flags().IntVar(&startTimestamp, "startTimestamp", env.ParseNumFromEnv(common.EnvVarStartTimestamp, common.DefaultStartTimestamp, 0, math.MaxInt64), "from timestamp")
	command.Flags().IntVar(&endTimestamp, "endTimestamp", env.ParseNumFromEnv(common.EnvVarEndTimestamp, common.DefaultEndTimestamp, 0, math.MaxInt64), "end timestamp")
	command.Flags().StringVar(&totalReward, "totalReward", env.StringFromEnv(common.EnvVarTotalReward, common.DefaultTotalReward), "total reward")
	command.MarkFlagRequired("cycle")

	return command
}
