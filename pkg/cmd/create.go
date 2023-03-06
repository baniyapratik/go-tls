package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-tls/pkg/key"
)

var keyOut string
var keyLength int

func init() {
	rootCmd.AddCommand(keyCreateCmd)
	keyCreateCmd.Flags().StringVarP(&keyOut, "key-out", "k", "key.pem", "destination path for key")
	keyCreateCmd.Flags().IntVarP(&keyLength, "key-length", "l", 4096, "key length")
}

var keyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create CA, certs, or keys",
	Long:  `command to create resources: ca, certs, keys`,
	Run: func(cmd *cobra.Command, args []string) {
		err := key.CreateRSAPrivateKeyAndSave(keyOut, keyLength)
		if err != nil {
			fmt.Printf("create key error: %s", err)
			return
		}
		fmt.Printf("key created %s with length %d", keyOut, keyLength)
	},
}
