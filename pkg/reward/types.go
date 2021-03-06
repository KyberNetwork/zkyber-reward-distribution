package reward

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type StrBigInt struct {
	*big.Int
}

func (sbi StrBigInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", sbi.String())), nil
}

func (sbi *StrBigInt) UnmarshalJSON(v []byte) error {
	if string(v) == "null" {
		return nil
	}

	if v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]
	}

	var z big.Int
	_, ok := z.SetString(string(v), 0)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", v)
	}

	sbi.Int = &z
	return nil
}

type OneReward struct {
	Index   int           `json:"index"`
	Amounts []*big.Int    `json:"amounts"`
	Proof   []common.Hash `json:"proof"`
}

func (or *OneReward) SetFromJsCompatibleVersion(jsor *jsCompatibleOneReward) {
	if or == nil {
		or = &OneReward{}
	}

	or.Index = jsor.Index
	or.Proof = jsor.Proof
	or.Amounts = []*big.Int{}
	for _, a := range jsor.Amounts {
		or.Amounts = append(
			or.Amounts,
			a.Int,
		)
	}
}

type jsCompatibleOneReward struct {
	Index   int           `json:"index"`
	Amounts []StrBigInt   `json:"amounts"`
	Proof   []common.Hash `json:"proof"`
}

func jsCompatibleOneRewardFromOneReward(or OneReward) *jsCompatibleOneReward {
	jsor := jsCompatibleOneReward{
		Index:   or.Index,
		Amounts: []StrBigInt{},
		Proof:   or.Proof,
	}
	for _, v := range or.Amounts {
		jsor.Amounts = append(
			jsor.Amounts,
			StrBigInt{big.NewInt(0).Set(v)},
		)
	}
	return &jsor
}

func (or OneReward) MarshalJSON() ([]byte, error) {
	jsor := jsCompatibleOneRewardFromOneReward(or)

	return json.Marshal(*jsor)
}

func (or *OneReward) UnmarshalJSON(v []byte) error {
	jsor := jsCompatibleOneReward{}
	err := json.Unmarshal(v, &jsor)
	if err != nil {
		return err
	}
	or.SetFromJsCompatibleVersion(&jsor)

	return nil
}

type Rewards struct {
	PhaseId        int                          `json:"phaseId"`
	Tokens         []common.Address             `json:"tokens"`
	Amounts        []*big.Int                   `json:"amounts"`
	StartTimestamp uint64                       `json:"startTimestamp"`
	EndTimestamp   uint64                       `json:"endTimestamp"`
	UserRewards    map[common.Address]OneReward `json:"userRewards"`
}
