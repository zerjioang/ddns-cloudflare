package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zerjioang/ddns-cloudflare/api"
)

const (
	version = "0.3.1"
)

var rootCmd = &cobra.Command{
	Use:   "ddns-cloudflare",
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
	Short: "update DDNS IP",
	Long:  `start IP address update procedure using Cloudllare v4 API`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := api.Start(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "start monitor",
	Long:  `start a foreground monitor that updates every 10 minutes the IP of the device`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := api.Monitor(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  `prints ddns-cloudflare agent application version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(monitorCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute is the main command line application entry point
func Execute() error {
	return rootCmd.Execute()
}
