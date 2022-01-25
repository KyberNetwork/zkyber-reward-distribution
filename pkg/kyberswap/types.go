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
	cfg *Cfg
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
