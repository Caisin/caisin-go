package list

import (
	"github.com/Caisin/caisin-go/tools/gvm/util"
	"github.com/Caisin/caisin-go/utils/lists"
	"github.com/Caisin/caisin-go/utils/strutil"
	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
	"strings"
)

var (
	update   = false //是否更新索引
	StartCmd = &cobra.Command{
		Use:          "idx",
		Short:        "go version idx",
		Example:      "gvm idx",
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

}

func run() error {
	doc, err := htmlquery.LoadURL("https://golang.google.cn/dl/")
	if err != nil {
		return err
	}
	nodes := htmlquery.Find(doc, `//div[contains(@class,"collapsed")]/h3[contains(@class,"toggleButton")]`)
	list := make([]string, 0)
	version := util.GetCurrentVersion()
	for _, node := range nodes {
		if node.FirstChild != nil {
			data := node.FirstChild.Data
			if strutil.IsNotBlank(data) {
				v := strings.ReplaceAll(data, " ▸", "")
				if v == version.Version {
					v = strutil.Green(v + " current")
				}
				list = append(list, v)
			}
		}
	}
	lists.PrintRev(list)
	return nil
}
