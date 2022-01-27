package kyberswap

import (
	"context"
	"fmt"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/util"
	"k8s.io/apimachinery/pkg/util/sets"
	"strconv"

	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/config"
	"github.com/machinebox/graphql"
)

func NewFetcher(cfg *Cfg, startTimestamp, endTimestamp uint64) (IRunner, error) {
	return &Fetcher{
		cfg:            cfg,
		startTimestamp: startTimestamp,
		endTimestamp:   endTimestamp,
	}, nil
}

func (s *Fetcher) Run() error {
	s.UpdateData(s.cfg.Common)

	return nil
}

func (s *Fetcher) UpdateData(cfg *config.Common) {
	ctx := context.Background()

	swapUsers, err := s.getSwapUsersList(ctx, cfg)
	if err != nil {
		fmt.Printf("failed to get router exchanges, err: %v", err)
	}
	fmt.Printf("got %d users who had swapped\n", len(swapUsers))

	addLiquidityUsers, err := s.getAddLiquidityUsersList(ctx, cfg)

	if err != nil {
		fmt.Printf("failed to get liquidity popsitions, err: %v", err)
	}
	fmt.Printf("got %d users who had added liquidity\n", len(addLiquidityUsers))

	totalUsers := util.SliceUnion(swapUsers, addLiquidityUsers)

	if err := util.WriteUsersListToFile(totalUsers, common.DataFolder, fmt.Sprintf(
		"%s/users_list_%d.json",
		common.DataFolder,
		s.cfg.Common.ChainID,
	)); err != nil {
		fmt.Printf("can not write data to file, err: %v", err)
	}
}

func (s *Fetcher) getSwapUsersList(ctx context.Context, cfg *config.Common) ([]string, error) {
	lastTimestamp := s.startTimestamp

	client := graphql.NewClient(cfg.Subgraph.Aggregator)

	usersSet := sets.NewString()

	mergeAddresses := func(addresses []string) {
		for _, a := range addresses {
			usersSet.Insert(a)
		}
	}

	for lastTimestamp <= s.endTimestamp {
		fmt.Printf("crawling router exchanges from timestamp: %d\n", lastTimestamp)
		query := fmt.Sprintf(routerLogQuery, graphFirstLimit, lastTimestamp, s.endTimestamp)
		req := graphql.NewRequest(query)

		var resp struct {
			Data []*RouterResp `json:"routerExchanges"`
		}

		if err := client.Run(ctx, req, &resp); err != nil {
			fmt.Printf("failed to query subgraph, err: %v\n", err)
			return nil, err
		}

		var usersArray []string

		for _, r := range resp.Data {
			usersArray = append(usersArray, r.UserAddress)
		}

		mergeAddresses(usersArray)

		if len(resp.Data) < graphFirstLimit {
			fmt.Println("no more router exchanges, stop crawling")
			break
		}

		t, err := strconv.ParseUint(resp.Data[len(resp.Data)-1].Timestamp, 10, 64)
		if err != nil {
			return nil, err
		}

		lastTimestamp = t
	}

	return usersSet.List(), nil
}

func (s *Fetcher) getAddLiquidityUsersList(ctx context.Context, cfg *config.Common) ([]string, error) {
	lastTimestamp := s.startTimestamp

	client := graphql.NewClient(cfg.Subgraph.Exchange)

	usersSet := sets.NewString()

	mergeAddresses := func(addresses []string) {
		for _, a := range addresses {
			usersSet.Insert(a)
		}
	}

	for lastTimestamp <= s.endTimestamp {
		fmt.Printf("crawling liquidity positions from timestamp: %d\n", lastTimestamp)
		query := fmt.Sprintf(addLiquidityQuery, graphFirstLimit, lastTimestamp, s.endTimestamp)
		req := graphql.NewRequest(query)

		var resp struct {
			Data []*LiquidityPositionResp `json:"liquidityPositionSnapshots"`
		}

		if err := client.Run(ctx, req, &resp); err != nil {
			fmt.Printf("failed to query subgraph, err: %v\n", err)
			return nil, err
		}

		var usersArray []string

		for _, r := range resp.Data {
			usersArray = append(usersArray, r.User.Id)
		}

		mergeAddresses(usersArray)

		if len(resp.Data) < graphFirstLimit {
			fmt.Println("no more router exchanges, stop crawling")
			break
		}

		t, err := strconv.ParseUint(resp.Data[len(resp.Data)-1].Timestamp, 10, 64)
		if err != nil {
			return nil, err
		}

		lastTimestamp = t
	}

	return usersSet.List(), nil
}
