package util

import (
	"fmt"
	"testing"
)

func TestUpdateVersionIndex(t *testing.T) {
	index, err := GetSetting()
	fmt.Println(index, err)
}
