# ZKyber reward distribution

## What does this tool do

This tool is built by Kyber team in order to provide an implementation to get the list of eligible users for ZKyber airdrop

With that in mind, Kyber team implemented the tool to do the following steps:

1. The tool reads from the subgraphs to get the list of users who swap or add liquidity in the selected timeframe.
2. The tool then reads the list of ZKyber eligible users (at the moment, it's the file ./data/zkyber_users_list.json).
3. The tool finds the users who belong to both lists.
4. The tool generates a new reward distribution content in a file which has the same format as this [example json](https://github.com/KyberNetwork/zkyber-reward-distribution/blob/main/results/latest_merkle_data.json
   ). 
5. The file provides all the information needed for a new reward distribution phase.

## How to use
1. Clone this repository to your local machine
2. Install Golang 1.17 or above
3. Install Node and yarn
4. Go to the the `merkle` folder, run `yarn install` to install the tool's dependency
5. Go back to the root folder and run `./compile_and_run.sh`

After a while if you see something like this:

```
==============================

Number of eligible users: 4533

==============================
Writing users list to ./results/phase_1/users_list.json...
Writing reward data to ./results/phase_1/reward_data.json...
Writing merkle data to ./results/phase_1/merkle_data.json ...
Done.
```

You have run the tool successfully and the reward distribution content is saved to `./results/phase_#{phaseId}/merkle_data.json`.

*Note: numbers in the output can change*.

## Global Environments

You can override any of the environment variables here. If you don't, the tool will get the default value from `pkg/common/default.go`

| Environment variable | Note                                                                              |
|----------------------|-----------------------------------------------------------------------------------|
| CHAIN_ID             | Log level (`trace`,`debug`,`info`,`warn`,`error`,`fatal`,`panic`) default: `info` |
| SUBGRAPH_EXCHANGE    | Exchange subgraph url                                                             |
| SUBGRAPH_AGGREGATOR  | Aggregator subgraph url                                                           |
| PHASE_ID             | Phase ID (1, 2, 3 ...), at the moment you have set it manually                    |
| START_TIMESTAMP      | Start timestamp to calculate eligible users                                       |
| END_TIMESTAMP        | End timestamp to calculate eligible users                                         |
| TOTAL_REWARD         | Total reward in KNC                                                               |


---

## Flags

Use for `fetcher` command only. It should point to a `.yaml` file inside `pkg/config` folder.

| CLI flag | Note        |
|----------|-------------|
| c,config | config file |

