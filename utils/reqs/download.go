package reqs

import (
	"errors"
	"fmt"
	"github.com/Caisin/caisin-go/utils/files"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadFile(url string, localPath string) error {
	var (
		fsize int64
		buf   = make([]byte, 32*1024)
	)
	if files.Exists(localPath) {
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

	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fsize", fsize)
	//创建文件
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	buffer, err := io.CopyBuffer(file, resp.Body, buf)
	fmt.Println(buffer, err)
	if err == nil {
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
	}
	return err
}
