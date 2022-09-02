package _switch

import (
	"github.com/Caisin/caisin-go/tools/gvm/util"
	"github.com/spf13/cobra"
)

var (
	version  = ""
	StartCmd = &cobra.Command{
		Use:          "switch",
		Short:        "switch version",
		Example:      "gvm switch v1.19",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "update go version list index")

}

func run() error {
	return util.SwitchVersion(version)
}
