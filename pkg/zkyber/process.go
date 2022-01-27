package zkyber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	"strings"

	"github.com/KyberNetwork/zkyber-reward-distribution/pkg/common"
)

func ParseResultFile() ([]string, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/zkyber_users_list.json", common.DataFolder))
	parsedAddresses := sets.NewString()

	if err != nil {
		return nil, err
	}

	var data []UserRecord

	err = json.Unmarshal(file, &data)

	if err != nil {
		return nil, err
	}

	for _, record := range data {
		a := strings.Replace(record.OwnerAddress, "\\", "0", 1)
		parsedAddresses.Insert(a)
	}

	return parsedAddresses.List(), nil
}
