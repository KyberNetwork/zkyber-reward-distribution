package config

type Common struct {
	ChainID  int
	Subgraph Subgraph
}

type Subgraph struct {
	Exchange   string
	Aggregator string
}
