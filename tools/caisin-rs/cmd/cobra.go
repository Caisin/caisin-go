package cmd

import (
	"errors"
	"fmt"
	"gitee.com/Caisin/caisin-go/utils/recovery"
	"gitee.com/Caisin/caisin-go/utils/strutil"
	"github.com/spf13/cobra"
	"os"
)

// 启动类
var rootCmd = &cobra.Command{
	Use:          "caisin-rs",
	Short:        "crs",
	SilenceUsage: true,
	Long:         `caisin-rs`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(strutil.Red("args can not be empty"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `rust util`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(StartCmd)
}

// Execute : apply commands
func Execute() {
	defer recovery.TryE()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error(), err)
		os.Exit(-1)
	}
}
