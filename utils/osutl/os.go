package osutl

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

//type Os string

const (
	Windows = "windows"
	Linux   = "linux"
	Darwin  = "darwin" //macos
	Freebsd = "freebsd"
)

func Home() (string, error) {
	current, err := user.Current()
	if nil == err {
		return strings.ReplaceAll(current.HomeDir, "\\", "/"), nil
	}
	// cross compile support
	switch runtime.GOOS {
	case "windows":
		return homeWindows()
	default:
		return homeUnix()
	}
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil

}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return strings.ReplaceAll(home, "\\", "/"), nil

}
