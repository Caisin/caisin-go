package util

import (
	"fmt"
	"testing"
)

func TestUpdateVersionIndex(t *testing.T) {
	index, err := UpdateVersionIndex()
	fmt.Println(index, err)
}
