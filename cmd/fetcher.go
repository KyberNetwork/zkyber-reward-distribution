package cmd

import (
	"fmt"
	"math"

	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/kyberswap"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/util/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile            string
	chainId            int
	subgraphExchange   string
	subgraphAggregator string
	startTimestamp     int
	endTimestamp       int
)

func newFetcherCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "fetcher",
		Short: "Start ZKyber reward fetcher",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := &kyberswap.Cfg{}

			if cfgFile != "" {
				if err := viper.Unmarshal(c); err != nil {
					return err
				}
			}

			if chainId != 0 {
				c.Common.ChainID = chainId
			}

			if subgraphExchange != "" {
				c.Common.Subgraph.Exchange = subgraphExchange
			}

			if subgraphAggregator != "" {
				c.Common.Subgraph.Aggregator = subgraphAggregator
			}

			fmt.Printf("config: %+v\n", c)

			fetcher, err := kyberswap.NewFetcher(c, uint64(startTimestamp), uint64(endTimestamp))

			if err != nil {
				return err
			}

			return fetcher.Run()
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	cmd.Flags().IntVar(&chainId, "chainId", env.ParseNumFromEnv(common.EnvVarChainId, 0, 0, math.MaxInt32), "chain ID")
	cmd.Flags().StringVar(&subgraphExchange, "subgraphExchange", env.StringFromEnv(common.EnvVarSubgraphExchange, ""), "exchange subgraph url")
	cmd.Flags().StringVar(&subgraphAggregator, "subgraphAggregator", env.StringFromEnv(common.EnvVarSubgraphAggregator, ""), "aggregator subgraph url")
	cmd.Flags().IntVar(&startTimestamp, "startTimestamp", env.ParseNumFromEnv(common.EnvVarStartTimestamp, common.DefaultStartTimestamp, 0, math.MaxInt64), "from timestamp")
	cmd.Flags().IntVar(&endTimestamp, "endTimestamp", env.ParseNumFromEnv(common.EnvVarEndTimestamp, common.DefaultEndTimestamp, 0, math.MaxInt64), "end timestamp")

	return cmd
}

func initConfig() {
	if cfgFile == "" {
		return
	}

	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}
