package util

import (
	"fmt"
	"os"
	"testing"
)

func TestUpdateVersionIndex(t *testing.T) {
	index, err := GetSetting()
	fmt.Println(index, err)
}

func TestGopath(t *testing.T) {

	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getenv("GOROOT"))
}

func TestSwitchVersion(t *testing.T) {
	SwitchVersion("go1.19")
}
