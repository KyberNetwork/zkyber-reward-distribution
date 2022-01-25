package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteUsersListToFile(usersSet []string, path string, fileName string) error {
	jsonData, err := json.MarshalIndent(usersSet, "", "  ")
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700) // Create your folder
	}

	fmt.Printf("Writing users list to ./%s...", fileName)

	return ioutil.WriteFile(fileName, jsonData, 0744)
}
