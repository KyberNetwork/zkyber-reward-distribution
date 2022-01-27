package main

import (
	"os"

	"github.com/KyberNetwork/zkyber-reward-distribution/cmd"
)

func main() {
	rootCmd := cmd.NewCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
