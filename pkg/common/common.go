package common

const (
	EnvVarChainId            = "CHAIN_ID"
	EnvVarSubgraphExchange   = "SUBGRAPH_EXCHANGE"
	EnvVarSubgraphAggregator = "SUBGRAPH_AGGREGATOR"
	EnvVarPhaseId            = "PHASE_ID"
	EnvVarStartTimestamp     = "START_TIMESTAMP"
	EnvVarEndTimestamp       = "END_TIMESTAMP"
	EnvVarTotalReward        = "TOTAL_REWARD"
)

var ChainId = map[string]int{
	"Ethereum":  1,
	"Polygon":   137,
	"BSC":       56,
	"Avalanche": 43114,
	"Fantom":    250,
	"Cronos":    25,
}
