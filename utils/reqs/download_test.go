package reqs

import "testing"

func TestDownload(t *testing.T) {
	DownloadFile("https://golang.google.cn/dl/go1.9.windows-amd64.zip", "go1.9.windows-amd64.zip")
}
