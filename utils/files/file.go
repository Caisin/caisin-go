package files

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func Remove(filePath string) error {
	return os.Remove(filePath)
}
func RemoveAll(filePath string) error {
	return os.RemoveAll(filePath)
}

func Is(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir returns true if path is a directory,
// or returns false when it's a file or not exist.
func IsDir(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func FileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := os.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	os.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(ctx, id, group, line)
	}
}

// GetFileSize 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// GetCurrentPath 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// FileSize 获取文件大小
func FileSize(f *os.File) (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

/*// GetFileSize 获取文件大小
func GetFileSize(f string) (int64, error) {
	stat, err := os.Stat(f)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}*/

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// Exists 检查文件是否存在
func Exists(src string) bool {
	return !IsNotExist(src)
}

// IsNotExist 检查文件是否不存在
func IsNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if IsNotExist(src) {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// OpenOrCreateFile 覆盖写入
func OpenOrCreateFile(src string) (*os.File, error) {
	IsNotExistMkDir(path.Dir(src))
	return os.Create(src)
}

// OpenAppendFile 追加写入
func OpenAppendFile(src string) (*os.File, error) {
	IsNotExistMkDir(path.Dir(src))
	return os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}

// ReadStr 文件读取
func ReadStr(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("Use ioutil.ReadFile to read a file:", result)
		return result, nil
	} else {
		return "", err
	}
}

func ReadInt64(filePath string) int64 {
	str, err := ReadStr(filePath)
	if err != nil {
		return 0
	}
	return cast.ToInt64(str)
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
