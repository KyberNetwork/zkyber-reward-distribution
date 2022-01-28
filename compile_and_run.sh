go build -o zkyber_reward &&
./zkyber_reward fetcher -c pkg/config/ethereum.yaml &&
./zkyber_reward fetcher -c pkg/config/polygon.yaml &&
./zkyber_reward fetcher -c pkg/config/bsc.yaml &&
./zkyber_reward fetcher -c pkg/config/avalanche.yaml &&
./zkyber_reward fetcher -c pkg/config/fantom.yaml &&
./zkyber_reward fetcher -c pkg/config/cronos.yaml &&
./zkyber_reward reward --phaseId 1 &&
node merkle/src/generateMerkleRoot.js -f ../../results/phase_1/reward_data.json
