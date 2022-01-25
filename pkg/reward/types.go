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
	Index             int              `json:"index"`
	Tokens            []common.Address `json:"tokens"`
	CumulativeAmounts []*big.Int       `json:"cumulativeAmounts"`
	Proof             []common.Hash    `json:"proof"`
}

func (or *OneReward) SetFromJsCompatibleVersion(jsor *jsCompatibleOneReward) {
	if or == nil {
		or = &OneReward{}
	}

	or.Index = jsor.Index
	or.Tokens = jsor.Tokens
	or.Proof = jsor.Proof
	or.CumulativeAmounts = []*big.Int{}
	for _, a := range jsor.CumulativeAmounts {
		or.CumulativeAmounts = append(
			or.CumulativeAmounts,
			a.Int,
		)
	}
}

type jsCompatibleOneReward struct {
	Index             int              `json:"index"`
	Tokens            []common.Address `json:"tokens"`
	CumulativeAmounts []StrBigInt      `json:"cumulativeAmounts"`
	Proof             []common.Hash    `json:"proof"`
}

func jsCompatibleOneRewardFromOneReward(or OneReward) *jsCompatibleOneReward {
	jsor := jsCompatibleOneReward{
		Index:             or.Index,
		Tokens:            or.Tokens,
		CumulativeAmounts: []StrBigInt{},
		Proof:             or.Proof,
	}
	for _, v := range or.CumulativeAmounts {
		jsor.CumulativeAmounts = append(
			jsor.CumulativeAmounts,
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
	Cycle          int                          `json:"cycle"`
	StartTimestamp uint64                       `json:"startTimestamp"`
	EndTimestamp   uint64                       `json:"endTimestamp"`
	UserRewards    map[common.Address]OneReward `json:"userRewards"`
	Tokens         []common.Address             `json:"tokens"`
	Amounts        []*big.Int                   `json:"amounts"`
}
