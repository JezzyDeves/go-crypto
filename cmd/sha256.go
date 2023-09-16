package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var filePath string

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Get the SHA256 checksum of a file",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		hash := sha256.Sum256(file)

		fmt.Println(hex.EncodeToString(hash[:]))
	},
}

func init() {
	rootCmd.AddCommand(sha256Cmd)
	sha256Cmd.Flags().StringVarP(&filePath, "file", "f", "", "The path to the file you want to hash")
	sha256Cmd.MarkFlagRequired("file")
}
