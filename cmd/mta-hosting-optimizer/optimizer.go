package main

import (
	"os"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "optimizer",
	Short: "Root command of the mta-hosting-optimizer project",
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utilities.ErrorLog("could not execute mta-hosting-optimizer", err)
		os.Exit(1)
	}
}
