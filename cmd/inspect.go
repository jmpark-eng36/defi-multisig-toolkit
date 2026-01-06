package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{Use: "inspect [addr]", Short: "Inspect Safe", Args: cobra.ExactArgs(1), RunE: func(c *cobra.Command, args []string) error { fmt.Println("inspecting", args[0]); return nil }}
func init() { rootCmd.AddCommand(inspectCmd) }
// v14
