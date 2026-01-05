package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rpcURL string
var rootCmd = &cobra.Command{Use: "msig", Short: "DeFi Multisig Toolkit v13"}
func Execute() { if err := rootCmd.Execute(); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) } }
func init() { rootCmd.PersistentFlags().StringVar(&rpcURL, "rpc", os.Getenv("RPC_URL"), "RPC endpoint") }
