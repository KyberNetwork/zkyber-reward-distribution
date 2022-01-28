package reward

import (
	"encoding/json"
	"errors"
	"fmt"
	zCommon "github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/util"
	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/zkyber"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	"log"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
)

type KyberReward struct {
	startTimestamp uint64
	endTimestamp   uint64
	rewardTokens   []common.Address
	rewardAmounts  []*big.Int
}

func NewKyberReward(startTimestamp, endTimestamp uint64, totalReward string) *KyberReward {
	totalRewardBI := new(big.Int)
	totalRewardBI, ok := totalRewardBI.SetString(totalReward, 10)

	if !ok {
		log.Fatalln("SetString: can't set total reward amount")
	}

	return &KyberReward{
		startTimestamp,
		endTimestamp,
		[]common.Address{
			common.HexToAddress("0xfe56d5892BDffC7BF58f2E84BE1b2C32D21C308b"), // KNC is only claimable on BSC
		},
		[]*big.Int{
			totalRewardBI,
		},
	}
}

func (s *KyberReward) getResultFiles(root string) ([]string, error) {
	var files []string

	chainFileNames := make(map[string]bool)

	for _, chainId := range zCommon.ChainId {
		chainFileNames[fmt.Sprintf("users_list_%d.json", chainId)] = true
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if _, ok := chainFileNames[info.Name()]; ok {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *KyberReward) CalculateRewardForUsers(phaseId int) error {
	result := &Rewards{
		PhaseId:        phaseId,
		Tokens:         s.rewardTokens,
		Amounts:        s.rewardAmounts,
		StartTimestamp: s.startTimestamp,
		EndTimestamp:   s.endTimestamp,
		UserRewards:    map[common.Address]OneReward{},
	}

	filePaths, err := s.getResultFiles("data")
	allAddresses := sets.NewString()

	if err != nil {
		return err
	}

	if len(filePaths) != len(zCommon.ChainId) {
		return errors.New(fmt.Sprintf("There are %d chains, but %d result files", len(zCommon.ChainId), len(filePaths)))
	}

	for _, path := range filePaths {
		file, _ := ioutil.ReadFile(path)

		var data []string

		_ = json.Unmarshal(file, &data)
		allAddresses.Insert(data...)
	}

	zkyberUsersList, err := zkyber.ParseResultFile()

	if err != nil {
		return err
	}

	finalList := util.SliceIntersection(allAddresses.List(), zkyberUsersList)

	numUsers := len(finalList)

	fmt.Println("==============================")
	fmt.Printf("\nNumber of eligible users: %d\n\n", numUsers)
	fmt.Println("==============================")

	for _, a := range finalList {
		address := common.HexToAddress(a)
		one := OneReward{
			Amounts: []*big.Int{},
		}
		for i, _ := range s.rewardTokens {
			amountFBig := util.NewFloat().SetInt(s.rewardAmounts[i])

			userAmountFBig := util.NewFloat().Quo(
				amountFBig,
				util.NewFloat().SetInt(big.NewInt(int64(numUsers))),
			)
			amount, _ := userAmountFBig.Int(nil)
			one.Amounts = append(
				one.Amounts,
				amount,
			)
		}
		result.UserRewards[address] = one
	}

	// Write final list
	if err := util.WriteUsersListToFile(finalList, fmt.Sprintf(
		"%s/phase_%d",
		zCommon.ResultsFolder,
		result.PhaseId,
	), fmt.Sprintf(
		"%s/phase_%d/users_list.json",
		zCommon.ResultsFolder,
		result.PhaseId,
	)); err != nil {
		fmt.Printf("can not write data to file, err: %v", err)

		return err
	}

	// Write reward data
	if err := WriteRewardDataToFile(result, fmt.Sprintf(
		"%s/phase_%d",
		zCommon.ResultsFolder,
		result.PhaseId,
	), fmt.Sprintf(
		"%s/phase_%d/reward_data.json",
		zCommon.ResultsFolder,
		result.PhaseId,
	)); err != nil {
		fmt.Printf("can not write data to file, err: %v", err)

		return err
	}

	return nil
}

func WriteRewardDataToFile(rewardData *Rewards, path string, fileName string) error {
	jsonData, err := json.MarshalIndent(rewardData, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("Writing reward data to ./%s...\n", fileName)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700) // Create your file
	}

	if err = ioutil.WriteFile(fileName, jsonData, 0744); err != nil {
		return err
	}

	return nil
}
