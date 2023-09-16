package cmd

import (
	"github.com/spf13/cobra"
)

// bcryptCmd represents the bcrypt command
var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "Root command for bcrypt",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(bcryptCmd)
}
