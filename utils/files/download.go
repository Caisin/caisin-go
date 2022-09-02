package files

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, localPath string) error {
	var (
		buf = make([]byte, 32*1024)
	)
	if Exists(localPath) {
		fmt.Println(localPath, "already download")
		return nil
	}
	tmpFilePath := localPath + ".download"
	fmt.Println(tmpFilePath)
	//创建一个http client
	client := new(http.Client)
	//client.Timeout = time.Second * 60 //设置超时时间
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	//创建文件
	file, err := OpenOrCreateFile(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	_, err = io.CopyBuffer(file, resp.Body, buf)
	if err == nil {
		err = file.Close()
		if err != nil {
			return err
		}
		err = os.Rename(tmpFilePath, localPath)
	}
	return err
}
