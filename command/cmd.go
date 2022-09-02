package command

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Run 运行命令
// dir:运行目录
// 命令名
// args 参数
func Run(dir, name string, args ...string) string {
	return string(RunBytes(dir, name, args...))
}
func RunWithErr(dir, name string, args ...string) (string, error) {
	msg, err := RunBytesWithErr(dir, name, args...)
	return string(msg), err
}
func RunCurWithErr(name string, args ...string) (string, error) {
	msg, err := RunBytesWithErr(".", name, args...)
	switch runtime.GOOS {
	case "windows":
		bytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(msg)
		return string(bytes), err
	default:
		return string(msg), err
	}
}

func RunCur(name string, args ...string) string {
	return string(RunBytes(".", name, args...))
}

func RunBytes(dir, name string, args ...string) []byte {
	msg, err := RunBytesWithErr(dir, name, args...)
	if err != nil {
		msg = append(msg, []byte(err.Error())...)
	}
	switch runtime.GOOS {
	case "windows":
		bytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(msg)
		return bytes
	default:
		return msg
	}
}
func RunBytesWithErr(dir, name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd.CombinedOutput()
}
func RunWithConsole(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CmdExist(dir, name string) bool {
	run := Run(dir, name)
	if strings.Contains(run, name) && strings.Contains(run, "not found") {
		return false
	}
	return true
}

func Exist(name string) bool {
	return CmdExist(".", name)
}
