package command

import (
	"os"
	"os/exec"
	"strings"
)

// Run 运行命令
// dir:运行目录
// 命令名
// args 参数
func Run(dir, name string, args ...string) string {
	return string(RunBytes(dir, name, args...))
}

func RunBytes(dir, name string, args ...string) []byte {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	msg, err := cmd.CombinedOutput()
	if err != nil {
		msg = append(msg, []byte(err.Error())...)
	}
	return msg
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
