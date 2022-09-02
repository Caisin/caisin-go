package files

import (
	"testing"
)

func TestTarGz(t *testing.T) {
	//err := Compress("test.tar.gz", "E:/code/go/caisin-go/tools")
	DeCompress("test.tar.gz", "out")
	//fmt.Println(err)
}
