package list

import (
	"fmt"
	"gitee.com/Caisin/caisin-go/utils/strutil"
	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
	"strings"
)

var (
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

}

func run() error {
	doc, err := htmlquery.LoadURL("https://golang.google.cn/dl/")
	if err != nil {
		return err
	}
	nodes := htmlquery.Find(doc, `//div[contains(@class,"collapsed")]/h3[contains(@class,"toggleButton")]`)
	for _, node := range nodes {
		if node.FirstChild != nil {
			data := node.FirstChild.Data
			if strutil.IsNotBlank(data) {
				fmt.Println(strings.ReplaceAll(data, " ▸", ""))
			}
		}
	}
	return nil
}
