package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var (
	input string
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Runs password through the bcrypt hash",
	Run: func(cmd *cobra.Command, args []string) {
		hash, _ := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

		fmt.Println("Hash: ", hex.EncodeToString(hash))
	},
}

func init() {
	bcryptCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&input, "input", "i", "", "The string input to be hashed")
}
