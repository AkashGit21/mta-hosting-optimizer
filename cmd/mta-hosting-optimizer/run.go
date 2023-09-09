package main

import (
	"github.com/AkashGit21/mta-hosting-optimizer/internal/server"
	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
	"github.com/spf13/cobra"
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Starts running the application server",
		Run: func(cmd *cobra.Command, args []string) {
			srv, err := server.NewServer()
			if err != nil {
				utilities.ErrorLog("Error getting new server:", err)
				return
			}

			server.StartServer(srv)
		},
	}

	rootCmd.AddCommand(runCmd)
}
