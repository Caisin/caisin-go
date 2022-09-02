package osutl

import (
	"fmt"
	"github.com/Caisin/caisin-go/command"
)

func setWindowsEnv(name, value string) error {
	cur, err := command.RunCurWithErr("setx", name, value)
	fmt.Println(cur)
	return err
}

func appendWindowsEnv(name, value string) error {
	cur, err := command.RunCurWithErr("setx", name, "%"+name+"%;"+value)
	fmt.Println(cur)
	return err
}
