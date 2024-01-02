package cmd

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "starts literalog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
