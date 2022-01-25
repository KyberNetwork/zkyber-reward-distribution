package cmd

import (
	"fmt"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/config"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/kyberswap"
	"math"

	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/util/env"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	chainId            int
	subgraphExchange   string
	subgraphAggregator string
	startTimestamp     int
	endTimestamp       int
)

func newFetcherCmd() *cobra.Command {

	var command = &cobra.Command{
		Use:   "fetcher",
		Short: "Start ZKyber reward fetcher",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := &kyberswap.Cfg{
				Common: &config.Common{
					ChainID: chainId,
					Subgraph: config.Subgraph{
						Exchange:   subgraphExchange,
						Aggregator: subgraphAggregator,
					},
				},
			}

			c.Common.ChainID = chainId
			c.Common.Subgraph.Exchange = subgraphExchange
			c.Common.Subgraph.Aggregator = subgraphAggregator

			fmt.Printf("config: %+v\n", c)

			fetcher, err := kyberswap.NewFetcher(c)

			if err != nil {
				return err
			}

			return fetcher.Run()
		},
	}

	command.Flags().IntVar(&chainId, "chainId", env.ParseNumFromEnv(common.EnvVarChainId, common.DefaultChainId, 0, math.MaxInt32), "chain ID")
	command.Flags().StringVar(&subgraphExchange, "subgraphExchange", env.StringFromEnv(common.EnvVarSubgraphExchange, ""), "exchange subgraph url")
	command.Flags().StringVar(&subgraphAggregator, "subgraphAggregator", env.StringFromEnv(common.EnvVarSubgraphAggregator, ""), "aggregator subgraph url")
	command.Flags().IntVar(&startTimestamp, "startTimestamp", env.ParseNumFromEnv(common.EnvVarStartTimestamp, common.DefaultStartTimestamp, 0, math.MaxInt64), "from timestamp")
	command.Flags().IntVar(&endTimestamp, "endTimestamp", env.ParseNumFromEnv(common.EnvVarEndTimestamp, common.DefaultEndTimestamp, 0, math.MaxInt64), "end timestamp")

	return command
}
