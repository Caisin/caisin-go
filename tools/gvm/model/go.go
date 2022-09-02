package model

type GoVersion struct {
	Version     string //版本
	FileName    string //文件名
	DownloadUrl string //下载地址
	Kind        string //类型
	Os          string //系统
	Arch        string //架构
	Size        string //大小
	Sha256      string //架构
}
