package cmd

import (
	"fmt"
	"github.com/Caisin/caisin-go/utils/files"
	"github.com/Caisin/caisin-go/utils/osutl"
	"github.com/Caisin/caisin-go/utils/strutil"
	"github.com/spf13/cobra"
	"path"
)

var (
	url      string
	choice   string
	StartCmd = &cobra.Command{
		Use:          "crates",
		Short:        "change crates registry",
		Example:      "caisin-rs creates -r caisin",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&url, "url", "u", "config/settings.admin.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&choice, "choice", "c", "caisin", "choice exists repository")

}

func run() error {
	home, err := osutl.Home()
	if err != nil {
		return err
	}
	targetFile := path.Join(home, ".cargo", "config")
	file, err := files.OpenOrCreateFile(targetFile)
	defer file.Close()
	if err != nil {
		return err
	}
	switch choice {
	case "caisin":
		file.WriteString(`[source.crates-io]
registry = "https://github.com/rust-lang/crates.io-index"
replace-with = 'caisin'

[source.caisin]
registry = "https://github.com/Caisin/crates.io-index"

`)
	case "reset":
		file.WriteString("")
	default:
		if strutil.IsNotBlank(url) {
			file.WriteString(fmt.Sprintf(`[source.crates-io]
registry = "%s"`, url))
		}
	}
	return nil
}
