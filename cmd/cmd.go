package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zerjioang/ddns-cloudflare/api"
)

var rootCmd = &cobra.Command{
	Use:   "cfagent",
	Short: "CloudFlare DDNS Agent",
	Long:  `CloudFlare DDNS Agent`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var startCmd = &cobra.Command{
	Use:   "update",
	Short: "start service",
	Long:  `start service`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := api.Start(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
func Execute() error {
	return rootCmd.Execute()
}
