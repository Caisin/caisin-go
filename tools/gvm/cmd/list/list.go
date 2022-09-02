package list

import (
	"github.com/Caisin/caisin-go/tools/gvm/util"
	"github.com/Caisin/caisin-go/utils/lists"
	"github.com/Caisin/caisin-go/utils/strutil"
	"github.com/spf13/cobra"
	"strings"
)

var (
	update   = false
	key      = ""
	StartCmd = &cobra.Command{
		Use:          "list",
		Short:        "list all go version",
		Example:      "gvm list",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().BoolVarP(&update, "update", "u", false, "update go version list index")
	StartCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "search key")
}

func run() error {
	if update {
		_, err := util.UpdateVersionIndex()
		if err != nil {
			return err
		}
	}
	version := util.GetCurrentVersion()
	setting, err := util.GetSetting()
	if err != nil {
		return err
	}
	list := make([]string, 0)
	for _, v := range setting.VersionList {
		if v == version.Version {
			v = strutil.Green(v + " current")
		}
		if strutil.IsNotBlank(key) {
			if !strings.Contains(v, key) {
				continue
			}
		}
		list = append(list, v)
	}
	lists.PrintRev(list)
	return nil
}
