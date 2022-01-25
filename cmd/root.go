package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	GitCommit = "unknown"
	BuildDate = "unknown"
	Version   = "unreleased"
)

func showVersion() {
	fmt.Printf("Version:\t %s\n", Version)
	fmt.Printf("Git commit:\t %s\n", GitCommit)
	fmt.Printf("Date:\t\t %s\n", BuildDate)
}

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:          "reward",
		Short:        "ZKyber Reward Distribution",
		Long:         "ZKyber Reward Distribution",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			version := cmd.Flag("version").Value.String()
			if version == "true" {
				showVersion()
			} else {
				_ = cmd.Help()
				os.Exit(1)
			}
		},
	}

	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version information and exit. This flag is only available at the global level.")

	// Add commands
	rootCmd.AddCommand(newFetcherCmd())
	rootCmd.AddCommand(newRewardCmd())

	return rootCmd
}
