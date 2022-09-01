package test

import (
	"fmt"
	"github.com/Caisin/caisin-go/command"
	"testing"
)

func TestFmt(t *testing.T) {
	//run := command.RunWithConsole(".", "gofmt", "-l", "-w", "-s", "../state")
	//fmt.Println(run)
	for _, s := range []string{"go", "fmt", "gofmt", "ffmpeg"} {
		exist := command.Exist(s)
		fmt.Printf("%s,%t\n", s, exist)
	}
}
