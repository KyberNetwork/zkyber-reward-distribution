package kyberswap

import "github.com/KyberNetwork/zkyber-reward-distribution/pkg/config"

type Cfg struct {
	Env    string
	Common *config.Common
}

type IRunner interface {
	Run() error
}

type Fetcher struct {
	cfg            *Cfg
	startTimestamp uint64
	endTimestamp   uint64
}

type RouterResp struct {
	Id          string `json:"id"`
	Pair        string `json:"pair"`
	Token       string `json:"token"`
	Amount      string `json:"amount"`
	UserAddress string `json:"userAddress"`
	Timestamp   string `json:"time"`
	BlockNumber string `json:"blockNumber"`
	Tx          string `json:"tx"`
}

type LiquidityPositionResp struct {
	User                  LiquidityPositionUser `json:"user"`
	LiquidityTokenBalance string                `json:"liquidityTokenBalance"`
	Timestamp             uint64                `json:"timestamp"`
}

type LiquidityPositionUser struct {
	Id string `json:"id"`
}
