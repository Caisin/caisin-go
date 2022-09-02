package util

import (
	"encoding/json"
	"github.com/Caisin/caisin-go/command"
	"github.com/Caisin/caisin-go/tools/gvm/consts"
	"github.com/Caisin/caisin-go/tools/gvm/model"
	"github.com/Caisin/caisin-go/utils/files"
	"github.com/Caisin/caisin-go/utils/osutl"
	"github.com/Caisin/caisin-go/utils/strutil"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"os"
	"path"
	"strings"
)

func GetCurrentVersion() *model.GoVersion {
	run := command.Run(".", "go", "version")
	split := strings.Split(run, " ")
	version := &model.GoVersion{}
	if len(split) == 4 {
		version.Version = split[2]
		osAr := split[3]
		osArSp := strings.Split(osAr, "/")
		version.Os, version.Arch = osArSp[0], osArSp[1]
	}
	return version
}

func UpdateVersionIndex() (*model.Setting, error) {
	doc, err := htmlquery.LoadURL(consts.IndexUrl)
	if err != nil {
		return nil, err
	}
	setting := &model.Setting{}
	nodes := htmlquery.Find(doc, `//div[contains(@class,"expanded")]/h3[contains(@class,"toggleButton")]`)
	ret := make(map[string][]model.GoVersion)
	for _, node := range nodes {
		version := GetVersionNameByVersionNode(node)
		setting.VersionList = append(setting.VersionList, version)
		trs := htmlquery.Find(node.Parent, "//tbody/tr")
		versions := make([]model.GoVersion, 0)
		for _, tr := range trs {
			tds := htmlquery.Find(tr, "//td")
			if len(tds) != 6 {
				continue
			}
			fileNameTd, kindTd, osTd, archTd, sizeTd, shaTd := tds[0], tds[1], tds[2], tds[3], tds[4], tds[5]
			fileNameA := htmlquery.FindOne(fileNameTd, "//a")
			fileName := fileNameA.FirstChild.Data
			downloadUrl := htmlquery.SelectAttr(fileNameA, "href")

			goVersion := model.GoVersion{
				Version:     version,
				FileName:    fileName,
				DownloadUrl: consts.IndexHost + downloadUrl,
			}
			if kindTd.FirstChild != nil {
				goVersion.Kind = kindTd.FirstChild.Data
			}
			if osTd.FirstChild != nil {
				goVersion.Os = osTd.FirstChild.Data
			}
			if archTd.FirstChild != nil {
				goVersion.Arch = archTd.FirstChild.Data
			}
			if sizeTd.FirstChild != nil {
				goVersion.Size = sizeTd.FirstChild.Data
			}
			if shaTd.FirstChild != nil && shaTd.FirstChild.FirstChild != nil {
				goVersion.Sha256 = shaTd.FirstChild.FirstChild.Data
			}
			versions = append(versions, goVersion)
		}
		ret[version] = versions
	}
	setting.Index = ret
	file, err := files.OpenOrCreateFile(getSettingFileName())
	if err == nil {
		defer file.Close()
		bytes, err := json.Marshal(setting)
		if err == nil {
			_, err = file.Write(bytes)
		}
	}
	return setting, err
}

func GetVersionNameByVersionNode(node *html.Node) string {
	sufix := " ▸"
	sufixLen := len(sufix)
	if node.FirstChild != nil {
		data := node.FirstChild.Data
		if strutil.IsNotBlank(data) {
			return data[:len(data)-sufixLen]
		}
	}
	return ""
}

func getSettingFileName() string {
	home, _ := osutl.Home()
	return path.Join(home, consts.GvmIdxSettingFile)
}
func GetSetting() (*model.Setting, error) {
	var setting *model.Setting
	settingFileName := getSettingFileName()
	if files.Exists(settingFileName) {
		str, err := os.ReadFile(settingFileName)
		if err != nil {
			return nil, err
		}
		setting = &model.Setting{}
		err = json.Unmarshal(str, setting)
		if err != nil {
			return nil, err
		}
		return setting, nil
	}
	return UpdateVersionIndex()
}
