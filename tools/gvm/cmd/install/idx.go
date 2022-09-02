package install

import (
	"errors"
	"fmt"
	"github.com/Caisin/caisin-go/tools/gvm/util"
	"github.com/Caisin/caisin-go/utils/files"
	"github.com/Caisin/caisin-go/utils/lists"
	"github.com/Caisin/caisin-go/utils/osutl"
	"github.com/Caisin/caisin-go/utils/strutil"
	"github.com/spf13/cobra"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	version  = ""
	remove   = false
	StartCmd = &cobra.Command{
		Use:          "install",
		Short:        "go install",
		Example:      "gvm install",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "install go version")
	StartCmd.PersistentFlags().BoolVarP(&remove, "remove", "r", false, "uninstall version")

}

func run() error {
	if strutil.IsBlank(version) {
		return errors.New("-v 参数 版本号不能为空")
	}
	if !strings.HasPrefix(version, "go") {
		version = "go" + version
	}
	setting, err := util.GetSetting()
	if err != nil {
		return err
	}
	versionDir := path.Join(setting.GvmPath, version)
	if remove {
		return os.RemoveAll(versionDir)
	}
	if files.Exists(versionDir) {
		return errors.New(version + " already install")
	}
	v, ok := setting.Index[version]
	if !ok {
		for _, vs := range setting.VersionList {
			if strings.Contains(vs, version) {
				fmt.Println(vs)
			}
		}
		return errors.New(fmt.Sprintf("version %s not found", version))
	}
	lists.Print(v)
	fileName := fmt.Sprintf("%s.%s-%s", version, runtime.GOOS, runtime.GOARCH)
	switch runtime.GOOS {
	case osutl.Windows:
		fileName = fileName + ".zip"
	case osutl.Darwin, osutl.Linux, osutl.Freebsd:
		fileName = fileName + ".tar.gz"
	default:
		return errors.New("un support os " + runtime.GOOS)
	}
	for _, goVersion := range v {
		if goVersion.FileName == fileName {
			downloadUrl := goVersion.DownloadUrl
			realFileName := path.Join(versionDir, fileName)
			files.DownloadFile(downloadUrl, realFileName)
			switch runtime.GOOS {
			case osutl.Windows:
				files.Unzip(realFileName, versionDir)
			case osutl.Darwin, osutl.Linux, osutl.Freebsd:
				files.DeCompress(realFileName, versionDir)
			}
			os.Remove(realFileName)
		}
	}
	//files.DownloadFile(v.)
	return nil
}
